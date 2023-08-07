package activities

import (
	"context"
	"log/slog"
)

func SendPush(ctx context.Context, mobileID string) error {
	slog.Info("sending push", "mobileID", mobileID)
	return nil
}
