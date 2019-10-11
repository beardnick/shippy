// Package main provides ...
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/beardnick/shippy/consignment-service/proto/consignment"

	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	cmd.Init()
	client := pb.NewShippingServiceClient("micro.srv.consignment", microclient.DefaultClient)
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	fmt.Println("before parse")
	consignment, err := parseFile(file)
	fmt.Println("after parse")
	if err != nil {
		log.Fatalf("unable to parse file: %v", err)
	}
	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("couldn't create consignment: %v", err)
	}
	log.Printf("Created : %v", r.Created)
	r, err = client.GetConsignment(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("failed to get consignment")
	}
	log.Printf("consignments:\n %v", r.Consignments)
}
