package main

import (
	"context"
	"fmt"

	"net"
	pro "github.com/yamyy123/netxd-customer/netxd_customer"

	controllers "github.com/yamyy123/netxd-customer-controller/controller"
	services "github.com/yamyy123/netxd-dal/services"
	config "github.com/yamyy123/netxd_config/config"
	constants "github.com/yamyy123/netxd_config/constants"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	CustomerCollection := config.GetCollection(client, "netxdb", "customer")
	controllers.CustomerService = services.InitCustomerService(CustomerCollection, context.Background())

}

func main() {
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	// fmt.Println("hi")
	lis, err := net.Listen("tcp", constants.Port)
	// fmt.Println("hii")

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
