package main

import (
	"github.com/lucaskatayama/learn-temporal/examples/sagas/saga"
	"github.com/lucaskatayama/learn-temporal/examples/sagas/service_a"
	"github.com/lucaskatayama/learn-temporal/examples/sagas/service_b"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal("failed to dial")
	}

	w := worker.New(c, "sagas", worker.Options{
		Identity: "saga-worker",
	})

	w.RegisterActivity(service_a.ActivityA)
	w.RegisterActivity(service_b.ActivityB)

	w.RegisterActivity(service_a.CompensateActivityA)
	w.RegisterActivity(service_b.CompensateActivityB)

	w.RegisterWorkflow(saga.Workflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}

}
