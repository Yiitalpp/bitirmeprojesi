package controllers

import (
	"net/http"
	"project/database"
	"project/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TokenRepo struct {
	Db *gorm.DB
}

type StatusToken struct {
	Status string
}

func NewTokenController() *TokenRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Token{})
	return &TokenRepo{Db: db}
}

// create Token
func (repository *TokenRepo) CreateToken(c *gin.Context) {
	var token models.Token
	c.BindJSON(&token)
	user, _ := c.Get("user")
	tokenObject, err := models.CreateToken(repository.Db, user.(models.User))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, tokenObject)
}

// get Token by id
func (repository *TokenRepo) GetToken(tokenID uint) (*models.Token, error) {
	var token models.Token
	err := repository.Db.First(&token, tokenID).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (repository *TokenRepo) GetTokenByTokenString(tokenString string) (*models.Token, error) {
	var token models.Token
	err := repository.Db.Where("token = ? AND starting_date <= ? AND ending_date >= ?", tokenString, time.Now(), time.Now()).First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}
