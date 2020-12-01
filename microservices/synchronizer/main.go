package main

import (
	"context"
	"log"
	"time"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"

	"github.com/kzmake/dapr-clock/constants"
	"github.com/kzmake/dapr-clock/microservices/synchronizer/handler"
)

var serviceAddress = ":3000"

func main() {
	s, err := daprd.NewService(serviceAddress)
	if err != nil {
		log.Fatalf("failed to start the server: %+v", err)
	}

	if err := s.AddBindingInvocationHandler(constants.ComponentSynchronizer, handler.Synchronize); err != nil {
		log.Fatalf("error adding binding handler: %+v", err)
	}

	// 起動から10秒後に初回の時間同期を実施
	go func() {
		time.Sleep(10 * time.Second)
		handler.Synchronize(context.Background(), &common.BindingEvent{})
	}()

	if err := s.Start(); err != nil {
		log.Fatalf("server error: %+v", err)
	}
}
