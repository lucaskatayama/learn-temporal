package workflows

import (
	"github.com/lucaskatayama/learn-temporal/examples/notification-hub/activities"
	"go.temporal.io/sdk/workflow"
	log "log/slog"
	"time"
)

type Event struct {
	ID     string
	Evt    string
	UserID string
}

func NotificationWorkflow(ctx workflow.Context, params Event) (err error) {
	var state string
	err = workflow.SetQueryHandler(ctx, "current_state", func() (string, error) {
		return state, nil
	})
	if err != nil {
		log.Error("setting query handler", "err", err)
		return
	}

	// configure activities
	activityOpts := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, activityOpts)

	// grab user information
	var user activities.User
	err = workflow.ExecuteActivity(ctx, activities.GetUserInformation, params.UserID).Get(ctx, &user)
	if err != nil {
		return
	}

	var doneEmail, doneMobile bool
	workflow.Go(ctx, func(ctx workflow.Context) {
		err := workflow.ExecuteActivity(ctx, activities.SendEmail, user.Email).Get(ctx, nil)
		if err != nil {
			log.Error("executing activity", "activity", "SendEmail", "err", err)
		}
		doneEmail = true
	})

	workflow.Go(ctx, func(ctx workflow.Context) {
		err := workflow.ExecuteActivity(ctx, activities.SendPush, user.MobileID).Get(ctx, nil)
		if err != nil {
			log.Error("executing activity", "activity", "SendPush", "err", err)
		}
		doneMobile = true
	})

	_ = workflow.Await(ctx, func() bool {
		return doneEmail && doneMobile
	})
	return
}
