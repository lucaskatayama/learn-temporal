package greeting

import (
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"time"
)

func Workflow(ctx workflow.Context, name string) error {

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    100 * time.Millisecond,
			BackoffCoefficient: 2,
			MaximumAttempts:    0,
		},
	})
	return workflow.ExecuteActivity(ctx, "greet", GreetParams{Greeting: "Hello from Golang", Name: name}).Get(ctx, nil)
}
