package handler

import (
	"todo/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
		auth.POST("/logout", h.logout)
		auth.POST("/send-verify-code", h.sendVerifyCode)
		auth.POST("/send-remove-code", h.sendRemoveCode)
	}

	// Simple group: v2
	user := router.Group("/user")
	{
		user.GET("/:id", h.getUserById)
		user.POST("/users", h.getUsers)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	return router
}
