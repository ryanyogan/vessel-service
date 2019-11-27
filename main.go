package main

import (
	"fmt"
	"log"
	"os"

	"github.com/micro/go-micro"
	pb "github.com/ryanyogan/vessel-service/proto/vessel"
)

const (
	defaultHost = "datastore:27017"
)

func createDummyData(repo Repository) {
	defer repo.Close()
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		repo.Create(v)
	}
}

func main() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = defaultHost
	}

	session, err := CreateSession(host)
	defer session.Close()

	if err != nil {
		log.Fatalf("Error connecting to the datastore: %v", err)
	}

	repo := &VesselRepository{session.Copy()}

	createDummyData(repo)

	server := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)
	server.Init()

	pb.RegisterVesselServiceHandler(server.Server(), &service{session})

	if err := server.Run(); err != nil {
		fmt.Println(err)
	}
}
