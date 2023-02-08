package storage

import (
	"github.com/iZarrios/gorm-psql-jwt-demo/models"
	"gorm.io/gorm"
)

// type UserDB = models.UserDB

type PostgresStore struct {
	DB *gorm.DB
}

var GlobalStore *PostgresStore

type User struct {
	Id       uint    `json:"-"`
	UserName *string `json:"username"`
	Password *string `json:"password"`
	Email    *string `json:"email"`
}

func (s *PostgresStore) CreateUser(user *User) error {
	err := s.DB.Create(&user).Error
	return err
}

func (s *PostgresStore) DeleteUser(id int) error {
	// use userDB
	err := s.DB.Delete(models.User{}, id).Error
	return err

}
func (s *PostgresStore) UpdateUser(user *models.User) error {
	//TODO
	// https://gorm.io/docs/update.html#Update-Changed-Fields
	return nil

}

func (s *PostgresStore) GetUsers() ([]models.User, error) {
	users := &[]models.User{}

	err := s.DB.Find(users).Error

	return *users, err
}

func (s *PostgresStore) GetUserByID(user *models.User) error {
	//TODO
	return nil

}

func (s *PostgresStore) GetUserByNumber(user *models.User) error {
	//TODO
	return nil
}
