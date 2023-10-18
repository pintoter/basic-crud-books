package v1

import (
	"github.com/pintoter/basic-crud-books/internal/service"

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

func (h *Handler) InitRoutes(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initBooksRoutes(v1)
		h.initUsersRoutes(v1)
	}
}
