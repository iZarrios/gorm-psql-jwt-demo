package storage

import (
	"github.com/iZarrios/gorm-psql-jwt-demo/models"
	"gorm.io/gorm"
)

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
	err := s.DB.Save(&user).Error
	return err
}

func (s *PostgresStore) GetUsers() ([]models.User, error) {
	users := &[]models.User{}

	err := s.DB.Find(users).Error

	return *users, err
}

// get user by id
func (s *PostgresStore) GetUserByID(id uint) (*models.User, error) {
	user := &models.User{Id: id}
	err := s.DB.First(&user).Error
	return user, err
}

// get user by username
func (s *PostgresStore) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{UserName: username}
	err := s.DB.First(&user).Error
	return user, err
}
