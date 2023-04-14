package handler

import (
	"github.com/Lunovoy/test-project-verba/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/health", h.checkHealth)

	api := router.Group("/api")
	{
		user := api.Group("/user")

		user.GET("/", h.addUser)
	}

	return router
}
