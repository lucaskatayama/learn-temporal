package activities

import (
	"context"
	"log/slog"
	"time"
)

func SendEmail(ctx context.Context, userEmail string) error {
	time.Sleep(3 * time.Second)
	slog.Info("sending email", "email", userEmail)
	return nil
}
