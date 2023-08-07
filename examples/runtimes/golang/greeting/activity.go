package greeting

import (
	"context"
	"fmt"
)

type GreetParams struct {
	Name     string `json:"name"`
	Greeting string `json:"greeting"`
}

func Greet(ctx context.Context, param GreetParams) (string, error) {
	return fmt.Sprintf("%s, %s!", param.Greeting, param.Name), nil
}
