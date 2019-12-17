package main

import pb "github.com/ryanyogan/vessel-service/proto/vessel"

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error) {

	}
}