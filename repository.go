package main

import (
	"context"

	pb "github.com/ryanyogan/vessel-service/proto/vessel"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
}

// VesselRepository holds the collection pointer to mongo
type VesselRepository struct {
	collection *mongo.Collection
}

// FindAvailable -- checks a spec against a map of vessels,
// if capacity and weight are below vessels, then return
// that vessel.
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	filter := bson.D{{
		"capacity",
		bson.D{{
			"$lte",
			spec.Capacity,
		}, {
			"$lte",
			spec.MaxWeight,
		}},
	}}
	var vessel *pb.Vessel
	if err := repo.collection.FindOne(context.TODO(), filter).Decode(&vessel); err != nil {
		return nil, err
	}

	return vessel, nil
}

// Create a new vessel
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	_, err := repo.collection.InsertOne(context.TODO(), vessel)
	return err
}
