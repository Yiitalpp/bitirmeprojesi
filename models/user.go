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

func CreateUser(db *gorm.DB, user *User) (err error) {
	err = db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
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

// get Users
func GetUsers(db *gorm.DB, users *[]User) (err error) {
	err = db.Find(users).Error
	if err != nil {
		return err
	}
	return nil
}

// get User by id
func GetUser(db *gorm.DB, user *User, id string) (err error) {
	err = db.Where("id = ?", id).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

// update a User
func UpdateUser(db *gorm.DB, user *User, id string) (err error) {
	err = db.Model(user).Where("id = ?", id).Updates(map[string]interface{}{"username": user.Username, "email": user.Email, "password": user.Password, "activation_code": user.ActivationCode, "active": user.Active, "last_login": user.LastLogin, "ip_address": user.IPAddress}).Error
	if err != nil {
		return err
	}
	return nil
}

// delete User
func DeleteUser(db *gorm.DB, user *User, id string) (err error) {
	err = db.Where("id = ?", id).Delete(user).Error
	if err != nil {
		return err
	}
	return nil
}
