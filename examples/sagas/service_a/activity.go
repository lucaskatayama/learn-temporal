package service_a

import (
	"context"
	"fmt"
	"log"
)

func ActivityA(ctx context.Context, name string) (string, error) {
	log.Println("Running A")
	return fmt.Sprintf("Hello %s", name), nil
}

func CompensateActivityA(ctx context.Context, name string) (string, error) {
	log.Println("Compensate A")
	return fmt.Sprintf("Compensate %s", name), nil
}
