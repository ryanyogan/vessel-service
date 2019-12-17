package main

import (
	"context"

	"github.com/ryanyogan/vessel-service/proto/vessel"
)

type handler struct {
	repository
}

// FindAvailable vessels
func (h *handler) FindAvailable(ctx context.Context, req *vessel.Specification, res *vessel.Response) error {
	vessel, err := h.repository.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

// Create a new vessel
func (h *handler) Create(ctx context.Context, req *vessel.Vessel, res *vessel.Response) error {
	if err := h.repository.Create(req); err != nil {
		return err
	}

	res.Vessel = req
	return nil
}
