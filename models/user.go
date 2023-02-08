package models

import (
	"time"

	"gorm.io/gorm"
)

// UserDB
type User struct {
	Id        uint           `gorm:"primary key;autoIncrement" json:"id"`
	UserName  string         `gorm:"not null" json:"username"`
	Password  string         `gorm:"not null" json:"password"`
	Email     string         `gorm:"not null;unique" json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type UserResponse struct {
	Id       uint   `json:"id"`
	UserName string `json:"username"`
	Password string ` json:"password"`
	Email    string `json:"email"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
