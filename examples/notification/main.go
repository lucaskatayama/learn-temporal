package main

import (
	"context"
	"github.com/lucaskatayama/learn-temporal/examples/notification/pkg/channels"
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

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, "HelloWorld", channels.Params{
		ID: time.Now().Format(time.RFC3339),
	})
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// Synchronously wait for the workflow completion.

	q, err := c.QueryWorkflow(context.Background(), "my id", we.GetRunID(), "current_state")
	if err != nil {
		log.Fatalln(err)
	}
	var result string
	err = q.Get(&result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
