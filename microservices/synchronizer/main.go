package main

import (
	"context"
	"log"
	"time"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"

	"github.com/kzmake/dapr-clock/microservices/synchronizer/infrastructure/pubsub"
	"github.com/kzmake/dapr-clock/microservices/synchronizer/interface/controller"
	"github.com/kzmake/dapr-clock/microservices/synchronizer/usecase/interactor"
)

var serviceAddress = ":3000"

func main() {
	s, err := pubsub.NewSyncService()
	if err != nil {
		panic(err)
	}
	i := interactor.NewSync(s)
	c := controller.NewSync(i)

	svc, err := daprd.NewService(serviceAddress)
	if err != nil {
		log.Fatalf("failed to start the server: %+v", err)
	}

	if err := svc.AddBindingInvocationHandler("synchronizer", c.Sync); err != nil {
		log.Fatalf("error adding binding handler: %+v", err)
	}

	// 起動から30秒後に初回の時間同期を実施
	go func() {
		time.Sleep(30 * time.Second)
		c.Sync(context.Background(), &common.BindingEvent{})
	}()

	if err := svc.Start(); err != nil {
		log.Fatalf("server error: %+v", err)
	}
}
