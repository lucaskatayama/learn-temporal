package service_b

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func ActivityB(ctx context.Context, name string) (string, error) {
	return "", errors.New("error")
}

func CompensateActivityB(ctx context.Context, name string) (string, error) {
	log.Println("Compensate B")
	return fmt.Sprintf("Compensating B %s", name), nil
}
