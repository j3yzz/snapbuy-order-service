package main

import (
	"fmt"
	"github.com/j3yzz/snapbuy-order-service/pkg/client"
	"github.com/j3yzz/snapbuy-order-service/pkg/config"
	"github.com/j3yzz/snapbuy-order-service/pkg/db"
	"github.com/j3yzz/snapbuy-order-service/pkg/pb"
	"github.com/j3yzz/snapbuy-order-service/pkg/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("failed to listing:", err)
	}

	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)

	fmt.Println("order service on:", c.Port)

	s := services.Server{
		H:          h,
		ProductSvc: productSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
