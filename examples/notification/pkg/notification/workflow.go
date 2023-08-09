package notification

import (
	"github.com/lucaskatayama/learn-temporal/examples/notification/pkg/channels"
	"github.com/lucaskatayama/learn-temporal/examples/notification/pkg/info"
	"go.temporal.io/sdk/workflow"
	"log"
	"time"
)

type Result struct {
	States []string
}

func Workflow(ctx workflow.Context, params channels.Params) (result Result, err error) {
	if err := workflow.SetQueryHandler(ctx, "current_state", func() (Result, error) {

		return result, nil
	}); err != nil {
		log.Fatal(err)
	}

	// configure workflow
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// execute getting user information activity
	p := channels.Params{
		ID: params.ID,
	}
	var state string
	if err := workflow.ExecuteActivity(ctx, info.User, p).Get(ctx, &state); err != nil {
		return result, err
	}

	// execute notifications
	notifiers := []any{
		channels.Email,
		channels.Mobile,
	}

	for _, notifier := range notifiers {
		var state string
		err := workflow.ExecuteActivity(ctx, notifier, p).Get(ctx, &state)
		if err != nil {

			return result, err
		}
		result.States = append(result.States, state)
	}
	return
}
