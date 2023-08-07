package main

import (
	"github.com/lucaskatayama/learn-temporal/examples/runtimes/golang/greeting"
	"go.temporal.io/sdk/activity"
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

	w := worker.New(c, "your-task-queue", worker.Options{
		Identity: "worker-golang",
	})

	w.RegisterWorkflowWithOptions(greeting.Workflow, workflow.RegisterOptions{
		Name:                          "YourWorkflow",
		DisableAlreadyRegisteredCheck: true,
	})
	w.RegisterActivityWithOptions(greeting.Greet, activity.RegisterOptions{
		Name: "greet",
	})

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
