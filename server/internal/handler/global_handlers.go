package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
