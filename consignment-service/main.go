package main

import (
	"context"
	"log"

	"github.com/beardnick/shippy/consignment-service/proto/consignment"
	vesselpb "github.com/beardnick/shippy/vessel-service/proto/vessel"

	micro "github.com/micro/go-micro"
)

var (
	port = ":50051"
)

// 持久层
type IRepository interface {
	Create(*consignment.Consignment) (*consignment.Consignment, error)
}

type Repository struct {
	consignments []*consignment.Consignment
}

func (repo *Repository) Create(consignment *consignment.Consignment) (*consignment.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

// 服务层

type service struct {
	repo         *Repository
	vesselClient vesselpb.VesselServiceClient
}

func (s *service) CreateConsignment(ctx context.Context, cons *consignment.Consignment, resp *consignment.Response) error {
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselpb.Specification{
		MaxWeight: cons.Weight,
		Capacity:  int32(len(cons.Containers)),
	})
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}
	cons.Vessel = vesselResponse.Vessel.Id
	_, err = s.repo.Create(cons)
	if err != nil {
		return err
	}
	// 注意这样写是错的
	// resp = &consignment.Response{Created: true, Consignment: cons}
	*resp = consignment.Response{Created: true, Consignment: cons}
	return nil
}

func (s *service) GetConsignment(ctx context.Context, req *consignment.GetRequest, resp *consignment.Response) error {
	*resp = consignment.Response{Consignments: s.repo.consignments}
	return nil
}

// go-micro的server
func main() {
	repo := &Repository{}
	// go-micro 使用服务名而非服务的端口来指明服务
	srv := micro.NewService(
		micro.Name("micro.srv.consignment"),
		micro.Version("latest"),
	)
	vesselClient := vesselpb.NewVesselServiceClient("micro.srv.vessel", srv.Client())
	srv.Init()
	consignment.RegisterShippingServiceHandler(srv.Server(), &service{repo, vesselClient})
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
