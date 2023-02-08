package storage

import "github.com/iZarrios/gorm-psql-jwt-demo/models"

type User = models.User


type Storage interface {

    CreateUser(*User) error
	DeleteUser(int) (int, error)
	UpdateUser(*User) error
	GetUsers() ([]*User, error)
	GetUserByID(int) (*User, error)
	GetUserByNumber(int) (*User, error)
}

