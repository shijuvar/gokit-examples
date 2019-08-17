package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/shijuvar/gokit-examples/services/account"
)

// Endpoints holds all Go kit endpoints for the Order service.
type Endpoints struct {
	CreateCustomer endpoint.Endpoint
	//AddMoneyToWallet endpoint.Endpoint
	//GetWalletBalance endpoint.Endpoint
	//MakePayment      endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the Account service.
func MakeEndpoints(s account.Service) Endpoints {
	return Endpoints{
		CreateCustomer: makeCreateCustomerEndpoint(s),
		//AddMoneyToWallet: makeAddMoneyToWalletEndpoint(s),
		//GetWalletBalance: makeGetWalletBalanceEndpoint(s),
		//MakePayment:      makeMakePaymentEndpoint(s),
	}
}

func makeCreateCustomerEndpoint(s account.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateCustomerRequest)
		err := s.CreateCustomer(ctx, req.Customer)
		return CreateCustomerResponse{Err: err}, nil
	}
}

//func makeAddMoneyToWalletEndpoint(s account.Service) endpoint.Endpoint {
//	return func(ctx context.Context, request interface{}) (interface{}, error) {
//		req := request.(AddMoneyToWalletRequest)
//		err := s.AddMoneyToWallet(ctx, req.CustomerID, req.Amount)
//		return AddMoneyToWalletResponse{Err: err}, nil
//	}
//}
//
//func makeGetWalletBalanceEndpoint(s account.Service) endpoint.Endpoint {
//	return func(ctx context.Context, request interface{}) (interface{}, error) {
//		req := request.(GetWalletBalanceRequest)
//		amount, err := s.GetWalletBalance(ctx, req.CustomerID)
//		return GetWalletBalanceResponse{Amount: amount, Err: err}, nil
//	}
//}
//
//func makeMakePaymentEndpoint(s account.Service) endpoint.Endpoint {
//	return func(ctx context.Context, request interface{}) (interface{}, error) {
//		req := request.(MakePaymentRequest)
//		err := s.MakePayment(ctx, req.CustomerID, req.Amount)
//		return MakePaymentResponse{Err: err}, nil
//	}
//}
