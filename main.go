package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/micro/go-micro"
	pb "github.com/ryanyogan/vessel-service/proto/vessel"
)

// Repository - is an interface that defines methods on this repo
type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

// VesselRepository -- Represents the repo (data-source)
type VesselRepository struct {
	vessels []*pb.Vessel
}

// FindAvailable -- checks a spec against a map of vessels,
// if capacity and weight are below vessels, then return
// that vessel.
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}

	return nil, errors.New("No vessel found by that spec")
}

// gRPC service handler
type service struct {
	repo repository
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty Boat", MaxWeight: 200000, Capacity: 500},
	}
	repo := &VesselRepository{vessels}

	srv := micro.NewService(
		micro.Name("transport.service.vessel"),
	)

	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
