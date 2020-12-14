package main

import (
	"log"

	daprd "github.com/dapr/go-sdk/service/grpc"

	"github.com/kzmake/dapr-clock/microservices/clock/infrastructure/hmshands"
	"github.com/kzmake/dapr-clock/microservices/clock/interface/controller"
	"github.com/kzmake/dapr-clock/microservices/clock/usecase/interactor"
)

var serviceAddress = ":3000"

func main() {
	r, err := hmshands.NewTimeRepository()
	if err != nil {
		panic(err)
	}
	i := interactor.NewNow(r)
	c := controller.NewClock(i)

	svc, err := daprd.NewService(serviceAddress)
	if err != nil {
		log.Fatalf("failed to start the server: %+v", err)
	}

	if err := svc.AddServiceInvocationHandler("now", c.Now); err != nil {
		log.Fatalf("error adding invocation handler: %+v", err)
	}

	if err := svc.Start(); err != nil {
		log.Fatalf("server error: %+v", err)
	}
}
