package saga

import (
	"context"
	"github.com/lucaskatayama/learn-temporal/examples/sagas/service_a"
	"github.com/lucaskatayama/learn-temporal/examples/sagas/service_b"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"time"
)

type Compensate func(ctx context.Context, name string) (string, error)

func Workflow(ctx workflow.Context, name string) (err error) {
	var state string
	var compensations []Compensate
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	})
	defer func() {
		if err != nil {
			for i := len(compensations) - 1; i >= 0; i-- {
				_ = workflow.ExecuteActivity(ctx, compensations[i], state).Get(ctx, &state)
			}
			err = nil
		}
	}()

	compensations = append(compensations, service_a.CompensateActivityA)
	err = workflow.ExecuteActivity(ctx, service_a.ActivityA, name).Get(ctx, &state)
	if err != nil {
		return
	}

	compensations = append(compensations, service_b.CompensateActivityB)
	err = workflow.ExecuteActivity(ctx, service_b.ActivityB, name).Get(ctx, &state)
	if err != nil {
		return
	}
	return nil
}
