package controllers

import (
	"errors"
	"net/http"
	"project/database"
	"project/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PlaneRepo struct {
	Db *gorm.DB
}

func NewPlaneController() *PlaneRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Plane{})
	return &PlaneRepo{Db: db}
}

func (repository *PlaneRepo) CreatePlane(c *gin.Context) {
	var plane models.Plane
	c.BindJSON(&plane)
	err := models.CreatePlane(repository.Db, &plane)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, plane)
}

func (repository *PlaneRepo) GetPlanes(c *gin.Context) {
	var planes []models.Plane
	err := models.GetPlanes(repository.Db, &planes)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, planes)
}

func (repository *PlaneRepo) GetPlane(c *gin.Context) {
	id := c.Param("id")
	var plane models.Plane
	err := models.GetPlane(repository.Db, &plane, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, plane)
}

func (repository *PlaneRepo) UpdatePlane(c *gin.Context) {
	id := c.Param("id")
	var plane models.Plane
	c.BindJSON(&plane)
	err := models.UpdatePlane(repository.Db, &plane, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, plane)
}

func (repository *PlaneRepo) DeletePlane(c *gin.Context) {
	id := c.Param("id")
	var plane models.Plane
	err := models.DeletePlane(repository.Db, &plane, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Plane deleted"})
}
