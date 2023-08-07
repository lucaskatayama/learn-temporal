package main

import (
	"github.com/lucaskatayama/learn-temporal/examples/notification/activities"
	"github.com/lucaskatayama/learn-temporal/examples/notification/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"log"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "notification", worker.Options{})

	w.RegisterWorkflowWithOptions(workflows.NotificationWorkflow, workflow.RegisterOptions{
		Name:                          "notification.Workflow",
		DisableAlreadyRegisteredCheck: true,
	})
	w.RegisterActivity(activities.GetUserInformation)
	w.RegisterActivity(activities.SendEmail)
	w.RegisterActivity(activities.SendPush)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
