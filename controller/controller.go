package controller

import (
	"context"
	pro "github.com/yamyy123/netxd-customer/netxd_customer"
	"github.com/yamyy123/netxd-dal/interfaces"
	"github.com/yamyy123/netxd-dal/models"
)

type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService interfaces.Icustomer
)

func (s *RPCServer) CreateCustomer(ctx context.Context, req *pro.CustomerRequest) (*pro.CustomerResponse, error) {
	dbCustomer := &models.CustomerRequest{Customer_Id: req.Customer_Id, CreatedAt: req.CreatedAt}
	result, err := CustomerService.CreateCustomer(dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			Customer_Id: result.Customer_Id,
			CreatedAt:   result.CreatedAt,
		}
		return responseCustomer, nil
	}
}
