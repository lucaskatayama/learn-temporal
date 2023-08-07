package notification

import (
	"github.com/lucaskatayama/learn-temporal/examples/notification/pkg/channels"
	"github.com/lucaskatayama/learn-temporal/examples/notification/pkg/info"
	"go.temporal.io/sdk/workflow"
	"log"
	"time"
)

type Result struct {
	State string
}

func Workflow(ctx workflow.Context, params channels.Params) (result Result, err error) {
	var res string
	if err := workflow.SetQueryHandler(ctx, "current_state", func() (Result, error) {

		return Result{State: res}, nil
	}); err != nil {
		log.Fatal(err)
	}

	// configure workflow
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 30 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// execute getting user information activity
	p := channels.Params{}
	if err := workflow.ExecuteActivity(ctx, info.User, p).Get(ctx, &res); err != nil {
		return result, err
	} else {
		p.ID = res
	}

	// execute notifications
	notifiers := []any{
		channels.Email,
		channels.Android,
		channels.IOs,
		channels.Web,
	}

	for _, notifier := range notifiers {
		err := workflow.ExecuteActivity(ctx, notifier, p).Get(ctx, &res)

		if err != nil {
			return Result{res}, err
		}
	}
	result.State = p.ID
	return
}
