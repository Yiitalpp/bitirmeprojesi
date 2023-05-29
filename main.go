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
	authMiddleware := controllers.AuthMiddleware(tokenRepo)

	userRepo := controllers.NewUserController()

	r.POST("/users", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:id", userRepo.GetUser)
	r.PUT("/users/:id", userRepo.UpdateUser)
	r.DELETE("/users/:id", userRepo.DeleteUser)

	r.POST("/register", userRepo.Register)
	r.POST("/login", userRepo.Login)
	r.POST("/logout", userRepo.Logout)

	ticketRepo := controllers.NewTicketController()
	r.POST("/tickets", ticketRepo.CreateTicket)
	r.GET("/tickets", ticketRepo.GetTickets)
	r.GET("/filtertickets", ticketRepo.FilterTickets)
	r.GET("/tickets/:id", ticketRepo.GetTicket)
	r.PUT("/tickets/:id", ticketRepo.UpdateTicket)
	r.DELETE("/tickets/:id", ticketRepo.DeleteTicket)

	bticketRepo := controllers.NewBTicketController()
	r.POST("/btickets", bticketRepo.CreateBTicket)
	r.GET("/btickets", bticketRepo.GetBTickets)
	r.GET("/btickets/:id", bticketRepo.GetBTicket)
	r.PUT("/btickets/:id", bticketRepo.UpdateBTicket)
	r.DELETE("/btickets/:id", bticketRepo.DeleteBTicket)

	planeRepo := controllers.NewPlaneController()
	r.POST("/planes", planeRepo.CreatePlane)
	r.GET("/planes", planeRepo.GetPlanes)
	r.GET("/planes/:id", planeRepo.GetPlane)
	r.PUT("/planes/:id", planeRepo.UpdatePlane)
	r.DELETE("/planes/:id", planeRepo.DeletePlane)

	//r.POST("/tickets/:ticket_id/book", userRepo.BookTicket)

	// Protected routes that require authentication
	protectedRoutes := r.Group("/")
	protectedRoutes.Use(authMiddleware)
	{
		protectedRoutes.POST("/tickets/:ticket_id/book", userRepo.BookTicket)
	}

	return r

}
