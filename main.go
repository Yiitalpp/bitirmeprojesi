package main

import (
	"net/http"
	"project/controllers"

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

	tokenRepo := controllers.NewTokenController()
	userRepo := controllers.NewUserController()
	r.POST("/register", userRepo.Register)
	r.POST("/login", userRepo.Login)
	r.POST("/logout", userRepo.Logout)

	return r

}
