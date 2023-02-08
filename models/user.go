package models

import (
	"gorm.io/gorm"
)

// UserDB
type User struct {
	Id       uint    `gorm:"primary key;autoIncrement" json:"_id"`
	UserName *string `gorm:"not null" json:"user_name"`
	Password *string `gorm:"not null" json:"password"`
	Email    *string `gorm:"not null;unique" json:"email"`
}

func MigrateUser(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	return err
}
