// Package main provides ...
package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/beardnick/shippy/vessel-service/proto/vessel"
	micro "github.com/micro/go-micro"
)

type IRepository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	vessels []*pb.Vessel
}

func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found for spec")
}

type service struct {
	repo IRepository
}

func (s *service) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	vessel, err := s.repo.FindAvailable(spec)
	if err != nil {
		return err
	}
	resp.Vessel = vessel
	return nil
}

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty Mc", MaxWeight: 20000, Capacity: 500},
	}
	repo := &VesselRepository{vessels}
	srv := micro.NewService(
		micro.Name("micro.srv.vessel-service"),
		micro.Version("latest"),
	)
	srv.Init()
	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
