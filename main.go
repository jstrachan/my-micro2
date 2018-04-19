package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/jstrachan/my-micro2/handler"
	"github.com/jstrachan/my-micro2/subscriber"

	example "github.com/jstrachan/my-micro2/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.my-micro2"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.my-micro2", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.my-micro2", service.Server(), subscriber.Handler)

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
