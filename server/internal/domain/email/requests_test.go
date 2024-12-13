package email

import (
	"errors"
	"testing"
)

func TestSendEmailRequest_Validate(t *testing.T) {
	type fields struct {
		Email   string
		Message string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "Valid email and message",
			fields: fields{
				Email:   "test@example.com",
				Message: "This is a valid message.",
			},
			wantErr: nil,
		},
		{
			name: "Invalid email format",
			fields: fields{
				Email:   "invalid-email",
				Message: "Message is here.",
			},
			wantErr: ErrInvalidEmailFormat,
		},
		{
			name: "Empty email",
			fields: fields{
				Email:   "",
				Message: "Message is here.",
			},
			wantErr: ErrInvalidEmailFormat,
		},
		{
			name: "Empty message",
			fields: fields{
				Email:   "test@example.com",
				Message: "",
			},
			wantErr: ErrInvalidEmailMessage,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &SendEmailRequest{
				Email:   tt.fields.Email,
				Message: tt.fields.Message,
			}
			err := r.Validate()
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
