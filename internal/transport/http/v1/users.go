package v1

import "github.com/gin-gonic/gin"

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	api.POST("/sign-up", h.signUp)
	api.POST("/sign-in", h.signIn)
}
