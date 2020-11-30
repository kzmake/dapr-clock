package main

import (
	"log"

	daprd "github.com/dapr/go-sdk/service/grpc"

	"github.com/kzmake/dapr-clock/constants"
	"github.com/kzmake/dapr-clock/microservices/synchronizer/handler"
)

var serviceAddress = ":3005"

func main() {
	s, err := daprd.NewService(serviceAddress)
	if err != nil {
		log.Fatalf("failed to start the server: %+v", err)
	}

	if err := s.AddBindingInvocationHandler(constants.ComponentSynchronizer, handler.Synchronize); err != nil {
		log.Fatalf("error adding binding handler: %+v", err)
	}

	if err := s.Start(); err != nil {
		log.Fatalf("server error: %+v", err)
	}
}
