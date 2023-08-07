package channels

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"log"
	"time"
)

func Web(ctx context.Context, params Params) (string, error) {
	go func() {
		activityInfo := activity.GetInfo(ctx)
		taskToken := activityInfo.TaskToken
		log.Println("email sent", params.ID)
		temporalClient, _ := client.Dial(client.Options{})

		// Complete the Activity.
		temporalClient.CompleteActivity(context.Background(), taskToken, fmt.Sprintf("[%s] email %s", params.ID, time.Now().Format(time.RFC3339Nano)), nil)
	}()

	return "", activity.ErrResultPending
}
