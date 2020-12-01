package main

import (
	"log"

	daprd "github.com/dapr/go-sdk/service/grpc"

	"github.com/kzmake/dapr-clock/constants"
	"github.com/kzmake/dapr-clock/microservices/clock/handler"
)

var serviceAddress = ":3000"

func main() {
	s, err := daprd.NewService(serviceAddress)
	if err != nil {
		log.Fatalf("failed to start the server: %+v", err)
	}

	if err := s.AddServiceInvocationHandler(constants.MethodNow, handler.Now); err != nil {
		log.Fatalf("error adding invocation handler: %+v", err)
	}

	if err := s.Start(); err != nil {
		log.Fatalf("server error: %+v", err)
	}
}
