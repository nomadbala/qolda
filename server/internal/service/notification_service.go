package service

import (
	"errors"
	"github.com/nomadbala/qolda/internal/domain/email"
	"github.com/resend/resend-go/v2"
)

var (
	ErrFailedSendEmail   = errors.New("failed while initializing notification service: failed to send email")
	ErrResendClientNull  = errors.New("failed while initializing notification service: resend client is null")
	ErrResendSenderNull  = errors.New("failed while initializing notification service: resend sender is null")
	ErrResendSubjectNull = errors.New("failed while initializing notification service: resend subject is null")
)

type NotificationService struct {
	Client  *resend.Client
	Sender  string
	Subject string
}

func NewNotificationService(client *resend.Client, sender, subject string) (*NotificationService, error) {
	if client == nil {
		return nil, ErrResendClientNull
	}

	if sender == "" {
		return nil, ErrResendSenderNull
	}

	if subject == "" {
		return nil, ErrResendSubjectNull
	}

	return &NotificationService{
		Client:  client,
		Sender:  sender,
		Subject: subject,
	}, nil
}

func (s *NotificationService) SendEmail(request email.SendEmailRequest) error {
	emailSendingParameters := &resend.SendEmailRequest{
		From:    s.Sender,
		To:      []string{request.Email},
		Subject: s.Subject,
		Html:    request.Message,
	}

	if _, err := s.Client.Emails.Send(emailSendingParameters); err != nil {
		return ErrFailedSendEmail
	}

	return nil
}
