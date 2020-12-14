package main

import (
	"log"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"

	"github.com/kzmake/dapr-clock/microservices/common/domain/event"

	"github.com/kzmake/dapr-clock/microservices/minute-hand/infrastructure/pubsub"
	"github.com/kzmake/dapr-clock/microservices/minute-hand/infrastructure/statestore"
	"github.com/kzmake/dapr-clock/microservices/minute-hand/interface/controller"
	"github.com/kzmake/dapr-clock/microservices/minute-hand/usecase/interactor"
)

var serviceAddress = ":3000"

func main() {
	r, err := statestore.NewHandRepository()
	if err != nil {
		panic(err)
	}
	s, err := pubsub.NewMovementService()
	if err != nil {
		panic(err)
	}
	in := interactor.NewNow(r)
	ii := interactor.NewIncrease(r, s)
	is := interactor.NewSync(r)
	c := controller.NewHand(in, ii, is)

	svc, err := daprd.NewService(serviceAddress)
	if err != nil {
		log.Fatalf("failed to start the server: %+v", err)
	}

	if err := svc.AddServiceInvocationHandler("now", c.Now); err != nil {
		log.Fatalf("error adding invocation handler: %+v", err)
	}

	if err := svc.AddTopicEventHandler(&common.Subscription{
		PubsubName: "pubsub",
		Topic:      event.Topic(event.Ticked60s{}),
		Route:      "/increase",
	}, c.Increase); err != nil {
		log.Fatalf("error adding event handler: %+v", err)
	}

	if err := svc.AddTopicEventHandler(&common.Subscription{
		PubsubName: "pubsub",
		Topic:      event.Topic(event.Synchronized{}),
		Route:      "/synchronize",
	}, c.Sync); err != nil {
		log.Fatalf("error adding event handler: %+v", err)
	}

	if err := svc.Start(); err != nil {
		log.Fatalf("server error: %+v", err)
	}
}
