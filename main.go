package main

import (
	"net/http"
	"project/controllers"
	"project/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	db := database.InitDb()
	authController := controllers.NewAuthController(db)

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)
	r.POST("/logout", authController.Logout)

	return r

}
