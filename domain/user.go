package domain

import (
	"errors"
	"time"
)

var (
	ErrInvalidEmail       = errors.New("invalid email format")                   // Email format is not valid
	ErrInvalidPassword    = errors.New("password must be at least 6 characters") // Password is too short
	ErrInvalidName        = errors.New("name cannot be empty")                   // Name field is required
	ErrUserNotFound       = errors.New("user not found")                         // User doesn't exist
	ErrUserAlreadyExists  = errors.New("user already exists")                    // User with this email already exists
	ErrInvalidCredentials = errors.New("invalid credentials")                    // Wrong email or password
	ErrUnauthorized       = errors.New("unauthorized access")                    // User not authorized
)

type User struct {
	Id        int64     `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"first_name" gorm:"not null"`
	LastName  string    `json:"last_name" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Password  string    `json:"password" gorm:"not null"`
	PhotoUrl  *string   `json:"photo_url"`
	Phone     *string   `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepositoryInterface interface {
	Create(user *User) error
	Update(user *User) error
	GetById(id int64) (*User, error)
	GetByEmail(email string) (*User, error)
	CheckUserExists(email string) (bool, error)
	List() ([]*User, error)
	Delete(id string) error
}

type UserServiceInterface interface {
	SignUp(user *User) error
	Login(email string, password string) (*User, error)
	List(userId int64) ([]*User, error)
	GenerateJWT(user *User) (string, error)
}
