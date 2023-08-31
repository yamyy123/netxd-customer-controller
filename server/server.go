package main

import (
	"context"
	"fmt"

	pro "netxd_project/netxd_customer"
	controllers "github.com/yamyy123/netxd-customer-controller/controller"
	"module/netxd_dal/service"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	CustomerCollection := config.GetCollection(client, "netxdb", "customer")
	controllers.CustomerService = service.InitCustomerService(CustomerCollection, context.Background())

}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pro.RegisterCustomerServiceServer(s, &controllers.RPCServer{})

	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}