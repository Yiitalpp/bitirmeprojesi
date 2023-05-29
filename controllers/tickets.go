package controllers

import (
	"errors"
	"net/http"
	"project/database"
	"project/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TicketRepo struct {
	Db *gorm.DB
}

func NewTicketController() *TicketRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Ticket{})
	return &TicketRepo{Db: db}
}

func (repository *TicketRepo) CreateTicket(c *gin.Context) {
	var ticket models.Ticket
	c.BindJSON(&ticket)
	err := models.CreateTicket(repository.Db, &ticket)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

func (repository *TicketRepo) GetTickets(c *gin.Context) {
	var tickets []models.Ticket
	err := models.GetTickets(repository.Db, &tickets)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, tickets)
}

func (repository *TicketRepo) FilterTickets(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	departureDateStr := c.Query("departureDate")
	returnDateStr := c.Query("returnDate")

	query := repository.Db.Model(&models.Ticket{})

	if from != "" {
		query = query.Where("`From` = ?", from)
	}

	if to != "" {
		query = query.Where("`To` = ?", to)
	}

	if departureDateStr != "" {
		departureDate, err := time.Parse("2006-01-02", departureDateStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid departureDate format"})
			return
		}
		departureDateFormatted := departureDate.Format("2006-01-02")
		query = query.Where("DATE(`departure_date`) = ?", departureDateFormatted)
	}

	if returnDateStr != "" {
		returnDate, err := time.Parse("2006-01-02", returnDateStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid returnDate format"})
			return
		}
		returnDateFormatted := returnDate.Format("2006-01-02")
		query = query.Where("DATE(`return_date`) = ?", returnDateFormatted)
	}

	var tickets []models.Ticket
	err := query.Find(&tickets).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

func (repository *TicketRepo) GetTicket(c *gin.Context) {
	id := c.Param("id")
	var ticket models.Ticket
	err := models.GetTicket(repository.Db, &ticket, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

func (repository *TicketRepo) UpdateTicket(c *gin.Context) {
	id := c.Param("id")
	var ticket models.Ticket
	c.BindJSON(&ticket)
	err := models.UpdateTicket(repository.Db, &ticket, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

func (repository *TicketRepo) DeleteTicket(c *gin.Context) {
	id := c.Param("id")
	var ticket models.Ticket
	err := models.GetTicket(repository.Db, &ticket, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = models.DeleteTicket(repository.Db, &ticket, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Ticket deleted"})
}
