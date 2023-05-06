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

	//tokenRepo := controllers.NewTokenController()
	userRepo := controllers.NewUserController()
	r.POST("/users", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:id", userRepo.GetUser)
	r.PUT("/users/:id", userRepo.UpdateUser)
	r.DELETE("/users/:id", userRepo.DeleteUser)
	r.POST("/register", userRepo.Register)
	r.POST("/login", userRepo.Login)
	r.POST("/logout", userRepo.Logout)

	return r

}
