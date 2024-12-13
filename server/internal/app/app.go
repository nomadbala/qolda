package app

import (
	"github.com/nomadbala/qolda/internal/config"
	"github.com/nomadbala/qolda/internal/handler"
	"github.com/nomadbala/qolda/internal/service"
	"github.com/nomadbala/qolda/pkg/server"
)

func Run() {
	configs, err := config.New()
	if err != nil {
		panic(err)
	}

	services, err := service.NewService(
		service.WithNotificationService(configs.Resend),
	)
	if err != nil {
		panic(err)
	}

	handlers, err := handler.NewHandler(
		handler.WithHTTPHandler(),
		handler.WithService(services),
	)

	servers := server.NewServer(handlers.Router)

	if err := servers.Run(); err != nil {
		panic(err)
	}
}
