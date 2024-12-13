package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nomadbala/qolda/internal/domain/email"
	"github.com/nomadbala/qolda/internal/service"
	"net/http"
)

func (h *Handler) SendEmailHandler(c *gin.Context) {
	var request email.SendEmailRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
	}

	if err := request.Validate(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
	}

	if err := h.Service.Notification.SendEmail(request); err != nil {
		if errors.Is(err, service.ErrFailedSendEmail) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		}
	}

	c.JSON(http.StatusOK, gin.H{})
}
