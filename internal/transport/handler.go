package transport

import (
	"books-app/internal/service"
	v1 "books-app/internal/transport/http/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	h.initAPI(router)

	return router
}

func (h *Handler) initAPI(e *gin.Engine) {
	handlerV1 := v1.NewHandler(h.service)
	api := e.Group("/api")
	{
		handlerV1.InitRoutes(api)
	}
}
