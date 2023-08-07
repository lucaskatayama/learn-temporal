package main

import (
	"github.com/lucaskatayama/learn-temporal/examples/notification/pkg/channels"
	"github.com/lucaskatayama/learn-temporal/examples/notification/pkg/info"
	"github.com/lucaskatayama/learn-temporal/examples/notification/pkg/notification"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
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

	w.RegisterWorkflow(notification.Workflow)
	//w.RegisterWorkflow(notify.Workflow)
	w.RegisterActivity(info.User)
	w.RegisterActivity(channels.Email)
	w.RegisterActivity(channels.Web)
	w.RegisterActivity(channels.Android)
	w.RegisterActivity(channels.IOs)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
