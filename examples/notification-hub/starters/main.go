package main

import (
	"context"
	"github.com/lucaskatayama/learn-temporal/examples/notification-hub/workflows"
	"go.temporal.io/sdk/client"
	"log"
	"time"
)

func main() {
	// The client is a heavyweight object that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        time.Now().Format(time.RFC3339),
		TaskQueue: "notification",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.NotificationWorkflow, workflows.Event{
		ID:     "123",
		Evt:    "Done something",
		UserID: "123",
	})
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

}
