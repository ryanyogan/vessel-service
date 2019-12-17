package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/ryanyogan/vessel-service/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func main() {
	srv := micro.NewService(
		micro.Name("transport.service.vessel"),
	)
	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	vesselCollection := client.Database("transport").Collection("vessel")
	repo := &VesselRepository{
		vesselCollection,
	}

	pb.RegisterVesselServiceHandler(srv.Server(), &handler{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
