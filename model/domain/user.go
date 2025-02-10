package domain

import (
	"time"
)

// sebuah model user yang mempresentasikan tabel di database.
type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
