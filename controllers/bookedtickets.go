package controllers

import (
	"errors"
	"net/http"
	"project/database"
	"project/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BTicketRepo struct {
	Db *gorm.DB
}

func NewBTicketController() *BTicketRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.BTicket{})
	return &BTicketRepo{Db: db}
}

func (repository *BTicketRepo) CreateBTicket(c *gin.Context) {
	var bTicket models.BTicket
	c.BindJSON(&bTicket)
	err := models.CreateBTicket(repository.Db, &bTicket)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, bTicket)
}

func (repository *BTicketRepo) GetBTickets(c *gin.Context) {
	var bTickets []models.BTicket
	err := models.GetBTickets(repository.Db, &bTickets)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, bTickets)
}

func (repository *BTicketRepo) GetBTicket(c *gin.Context) {
	id := c.Param("id")
	var bTicket models.BTicket
	err := models.GetBTicket(repository.Db, &bTicket, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, bTicket)
}

func (repository *BTicketRepo) UpdateBTicket(c *gin.Context) {
	id := c.Param("id")
	var bTicket models.BTicket
	c.BindJSON(&bTicket)
	err := models.UpdateBTicket(repository.Db, &bTicket, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, bTicket)
}

func (repository *BTicketRepo) DeleteBTicket(c *gin.Context) {
	id := c.Param("id")
	var bTicket models.BTicket
	err := models.GetBTicket(repository.Db, &bTicket, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = models.DeleteBTicket(repository.Db, &bTicket, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Booked Ticket deleted"})
}
