package main

import (
	"github.com/lucaskatayama/learn-temporal/examples/notification/pkg/channels"
	"github.com/lucaskatayama/learn-temporal/examples/notification/pkg/info"
	"github.com/lucaskatayama/learn-temporal/examples/notification/pkg/notification"
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

	w.RegisterWorkflowWithOptions(notification.Workflow, workflow.RegisterOptions{
		Name:                          "notification.Workflow",
		DisableAlreadyRegisteredCheck: true,
	})
	w.RegisterActivity(info.User)
	w.RegisterActivity(channels.Email)
	w.RegisterActivity(channels.Mobile)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
