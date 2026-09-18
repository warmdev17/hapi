package transport

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
}
