package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nomadbala/qolda/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	ErrFailedInitializeHandlerServiceNull = errors.New("failed initialize handler: services is null")
)

type Configuration func(h *Handler) error

type Handler struct {
	Router  *gin.Engine
	Service *service.Service
}

func NewHandler(configs ...Configuration) (*Handler, error) {
	handler := &Handler{}

	for _, config := range configs {
		if err := config(handler); err != nil {
			return nil, err
		}
	}

	return handler, nil
}

func WithHTTPHandler() Configuration {
	return func(h *Handler) error {
		h.Router = gin.Default()

		h.Router.GET("/health", h.HealthCheckHandler)

		h.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		apiHandlers := h.Router.Group("/api/v1")

		notificationHandlers := apiHandlers.Group("/email")
		notificationHandlers.POST("/send", h.SendEmailHandler)

		return nil
	}
}

func WithService(service *service.Service) Configuration {
	return func(h *Handler) error {
		if service == nil {
			return ErrFailedInitializeHandlerServiceNull
		}

		h.Service = service

		return nil
	}
}
