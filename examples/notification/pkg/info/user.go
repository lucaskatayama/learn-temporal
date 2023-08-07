package info

import (
	"context"
	"github.com/lucaskatayama/learn-temporal/examples/notification/pkg/channels"
)

func User(ctx context.Context, params channels.Params) (string, error) {
	return "123", nil
}
