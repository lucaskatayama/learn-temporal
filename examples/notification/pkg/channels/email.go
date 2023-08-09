package channels

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"log"
	"time"
)

func Email(ctx context.Context, params Params) (string, error) {
	go func() {
		activityInfo := activity.GetInfo(ctx)
		taskToken := activityInfo.TaskToken
		log.Println("email sent")
		temporalClient, _ := client.Dial(client.Options{})
		time.Sleep(1 * time.Second)
		// Complete the Activity.
		temporalClient.CompleteActivity(context.Background(), taskToken, fmt.Sprintf("email %s", time.Now().Format(time.RFC3339Nano)), nil)
	}()

	return "", activity.ErrResultPending
}
