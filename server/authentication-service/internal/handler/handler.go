package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrFailedInitializeHandler = errors.New("failed to initialize handler")
)

type Configuration func(h *Handler) error

type Handler struct {
	Router *gin.Engine
}

func NewHandler(configs ...Configuration) (*Handler, error) {
	handler := &Handler{}

	for _, config := range configs {
		if err := config(handler); err != nil {
			return nil, ErrFailedInitializeHandler
		}
	}

	return handler, nil
}

func WithHTTPHandler() Configuration {
	return func(h *Handler) error {
		h.Router = gin.Default()

		return nil
	}
}
