package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             uint   `json:"id" gorm:"primary_key"`
	Username       string `json:"username" gorm:"unique"`
	Email          string `json:"email" gorm:"unique"`
	Password       string `json:"password"`
	PlainPassword  string `gorm:"-"`
	ActivationCode string
	Active         bool
	LastLogin      *time.Time
	IPAddress      string
	CreatedAt      *time.Time
}

func Register(db *gorm.DB, user *User) (err error) {
	err = db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func Login(db *gorm.DB, user *User, email string) (err error) {
	err = db.Where("email = ?", email).First(user).Error
	if err != nil {
		return err
	}
	return nil
}
