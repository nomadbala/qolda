package service

import (
	"errors"
	"github.com/nomadbala/qolda/internal/config"
	"github.com/resend/resend-go/v2"
)

var (
	ErrFailedInitializeService = errors.New("failed to initialize service")
)

const (
	ResendSender  = "qolda@kiteo.app"
	ResendSubject = "Qolda"
)

type Configuration func(s *Service) error

type Service struct {
	Notification *NotificationService
}

func NewService(configs ...Configuration) (*Service, error) {
	service := &Service{}

	for _, config := range configs {
		if err := config(service); err != nil {
			return nil, ErrFailedInitializeService
		}
	}

	return service, nil
}

func WithNotificationService(config config.ResendConfig) Configuration {
	return func(s *Service) (err error) {

		s.Notification, err = NewNotificationService(
			resend.NewClient(config.APIKey),
			ResendSender,
			ResendSubject,
		)

		return err
	}
}
