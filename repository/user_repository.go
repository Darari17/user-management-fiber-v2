package repository

import (
	"database/sql"
	"errors"

	"github.com/Darari17/user-management/fiber/v2/model/domain"
)

type UserRepositoryImpl struct {
	db *sql.DB
}

type UserRepository interface {
	CreateRepo(user domain.User) (domain.User, error)
	UpdateRepo(user domain.User) (domain.User, error)
	DeleteRepo(id int) error
	GetRepo() ([]domain.User, error)
	FindByIdRepo(id int) (domain.User, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (u *UserRepositoryImpl) CreateRepo(user domain.User) (domain.User, error) {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	result, err := u.db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return domain.User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.User{}, err
	}

	query = "SELECT created_at FROM users WHERE id = ?"
	err = u.db.QueryRow(query, id).Scan(&user.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}

	user.ID = int(id)
	return user, nil
}

func (u *UserRepositoryImpl) UpdateRepo(user domain.User) (domain.User, error) {
	query := "UPDATE users SET username = ?, email = ?, password = ? WHERE id = ?"
	result, err := u.db.Exec(query, user.Username, user.Email, user.Password, user.ID)
	if err != nil {
		return domain.User{}, err
	}

	query = "SELECT updated_at FROM users WHERE id = ?"
	err = u.db.QueryRow(query, user.ID).Scan(&user.UpdatedAt)
	if err != nil {
		return domain.User{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return domain.User{}, err
	}

	if rowsAffected == 0 {
		return domain.User{}, errors.New("no user updated")
	}

	return user, nil
}

func (u *UserRepositoryImpl) DeleteRepo(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := u.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no user deleted")
	}

	return nil
}

func (u *UserRepositoryImpl) GetRepo() ([]domain.User, error) {
	query := "SELECT id, username, email, created_at, updated_at FROM users"
	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepositoryImpl) FindByIdRepo(id int) (domain.User, error) {
	var user domain.User
	query := "SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = ?"
	err := u.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, errors.New("user not found")
		}
		return domain.User{}, err
	}
	return user, nil
}
