package email

import (
	"errors"
	"net/mail"
)

var (
	ErrInvalidEmailFormat  = errors.New("send email request error: invalid email format")
	ErrInvalidEmailMessage = errors.New("send email request error: invalid email message")
)

type SendEmailRequest struct {
	Email   string `json:"email"   binding:"required"`
	Message string `json:"message" binding:"required"`
}

func (r *SendEmailRequest) Validate() error {
	if _, err := mail.ParseAddress(r.Email); err != nil {
		return ErrInvalidEmailFormat
	}

	if r.Message == "" {
		return ErrInvalidEmailMessage
	}

	return nil
}
