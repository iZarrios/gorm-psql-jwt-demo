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


// func (s *PostgresStore) CreateUser(user *User) error {
func (s *PostgresStore) CreateUser(user *models.User) error {
	err := s.DB.Create(&user).Error
	return err
}

func (s *PostgresStore) DeleteUser(id uint) error {

	user := &models.User{Id: id}
	err := s.DB.First(&user).Error

	// Delete results in a sucess anyways?
	// err := s.DB.Delete(&user).Error

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
