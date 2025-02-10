package service

import (
	"github.com/Darari17/user-management/fiber/v2/model/domain"
	"github.com/Darari17/user-management/fiber/v2/model/dto"
	"github.com/Darari17/user-management/fiber/v2/repository"
	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	Repo     repository.UserRepository
	Validate *validator.Validate
}

type UserService interface {
	CreateService(request dto.CreateRequest) (dto.Response, error)
	UpdateService(request dto.UpdateRequest) (dto.Response, error)
	DeleteService(id int) error
	GetService() ([]dto.Response, error)
	FindByIdService(id int) (dto.Response, error)
}

func NewUserService(repo repository.UserRepository) UserService {
	return &UserServiceImpl{Repo: repo, Validate: validator.New()}
}

func (u *UserServiceImpl) CreateService(request dto.CreateRequest) (dto.Response, error) {
	if err := u.Validate.Struct(request); err != nil {
		return dto.Response{}, err
	}

	user := domain.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	create, err := u.Repo.CreateRepo(user)
	if err != nil {
		return dto.Response{}, err
	}

	response := dto.Response{
		ID:        create.ID,
		Username:  create.Username,
		Email:     create.Email,
		CreatedAt: create.CreatedAt,
		UpdatedAt: create.UpdatedAt,
	}

	return response, nil

}

func (u *UserServiceImpl) UpdateService(request dto.UpdateRequest) (dto.Response, error) {
	if err := u.Validate.Struct(request); err != nil {
		return dto.Response{}, err
	}

	user, err := u.Repo.FindByIdRepo(request.ID)
	if err != nil {
		return dto.Response{}, err
	}

	requestUpdate := domain.User{
		ID:       request.ID,
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	if requestUpdate.Username == "" {
		requestUpdate.Username = user.Username
	}
	if requestUpdate.Email == "" {
		requestUpdate.Email = user.Email
	}
	if requestUpdate.Password == "" {
		requestUpdate.Password = user.Password
	}

	commitUpdate, err := u.Repo.UpdateRepo(requestUpdate)
	if err != nil {
		return dto.Response{}, err
	}

	response := dto.Response{
		ID:        commitUpdate.ID,
		Username:  commitUpdate.Username,
		Email:     commitUpdate.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: commitUpdate.UpdatedAt,
	}

	return response, nil
}

func (u *UserServiceImpl) DeleteService(id int) error {
	user, err := u.Repo.FindByIdRepo(id)
	if err != nil {
		return err
	}

	return u.Repo.DeleteRepo(user.ID)
}

func (u *UserServiceImpl) GetService() ([]dto.Response, error) {
	users, err := u.Repo.GetRepo()
	if err != nil {
		return nil, err
	}

	var responses []dto.Response
	for _, user := range users {
		responses = append(responses, dto.Response{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return responses, nil
}

func (u *UserServiceImpl) FindByIdService(id int) (dto.Response, error) {
	user, err := u.Repo.FindByIdRepo(id)
	if err != nil {
		return dto.Response{}, err
	}

	response := dto.Response{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}
