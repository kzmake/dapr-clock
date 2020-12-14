package main

import (
	"log"

	daprd "github.com/dapr/go-sdk/service/grpc"

	"github.com/kzmake/dapr-clock/microservices/ticker/infrastructure/pubsub"
	"github.com/kzmake/dapr-clock/microservices/ticker/interface/controller"
	"github.com/kzmake/dapr-clock/microservices/ticker/usecase/interactor"
)

var serviceAddress = ":3000"

func main() {
	s, err := pubsub.NewMovementService()
	if err != nil {
		panic(err)
	}
	i := interactor.NewTick(s)
	c := controller.NewTick(i)

	svc, err := daprd.NewService(serviceAddress)
	if err != nil {
		log.Fatalf("failed to start the server: %+v", err)
	}

	// FIXME: コンポーネント: ticker を定義し、 1sec毎に処理するためのinput bindingを追加する
	_ = c.Tick

	if err := svc.Start(); err != nil {
		log.Fatalf("server error: %+v", err)
	}
}
