package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	UserID       uint
	Token        string
	StartingDate *time.Time
	EndingDate   *time.Time
}

// create a Token
func CreateToken(db *gorm.DB, user User) (Token, error) {
	var token Token

	token.Token = uuid.New().String()
	token.UserID = user.ID
	startingDate := time.Now()
	token.StartingDate = &startingDate
	endDate := startingDate.Add(time.Hour * 24 * 30) // Token is valid for 30 days
	token.EndingDate = &endDate

	if err := db.Create(&token).Error; err != nil {
		return token, err
	}
	return token, nil
}

// get Token by id
func GetToken(db *gorm.DB, Token *Token, id string) (err error) {
	err = db.Where("id = ?", id).First(Token).Error
	if err != nil {
		return err
	}
	return nil
}
