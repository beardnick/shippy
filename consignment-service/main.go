package main

import (
	"context"
	"log"
	"shippy/consignment-service/proto/consignment"

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
	repo *Repository
}

func (s *service) CreateConsignment(ctx context.Context, cons *consignment.Consignment, resp *consignment.Response) error {
	_, err := s.repo.Create(cons)
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
	srv.Init()
	consignment.RegisterShippingServiceHandler(srv.Server(), &service{repo})
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
