package main

import (
	"context"
	"log"
	"net"
	"shippy/consignment-service/proto/consignment"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

func (s *service) CreateConsignment(ctx context.Context, cons *consignment.Consignment) (*consignment.Response, error) {
	_, err := s.repo.Create(cons)
	if err != nil {
		return nil, err
	}
	return &consignment.Response{Created: true, Consignment: cons}, nil
}

func (s *service) GetConsignment(ctx context.Context, req *consignment.GetRequest) (*consignment.Response, error) {
	return &consignment.Response{Consignments: s.repo.consignments}, nil
}

// 这是实现了以grpc的server
func main() {
	repo := &Repository{}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen", err)
	}
	s := grpc.NewServer()
	consignment.RegisterShippingServiceServer(s, &service{repo})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
