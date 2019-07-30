package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"github.com/shijuvar/gokit-examples/services/order"
)

// Endpoints holds all Go kit endpoints for the Order service.
type Endpoints struct {
	Create       endpoint.Endpoint
	GetByID      endpoint.Endpoint
	ChangeStatus endpoint.Endpoint
}

// MakeEndpoints initializes all Go kit endpoints for the Order service.
func MakeEndpoints(s order.Service) Endpoints {
	return Endpoints{
		Create:       makeCreateEndpoint(s),
		GetByID:      makeGetByIDEndpoint(s),
		ChangeStatus: makeChangeStatusEndpoint(s),
	}
}

func makeCreateEndpoint(s order.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest) // type assertion
		id, err := s.Create(ctx, req.Order)
		return CreateResponse{ID: id, Err: err}, nil
	}
}

func makeGetByIDEndpoint(s order.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)
		orderRes, err := s.GetByID(ctx, req.ID)
		return GetByIDResponse{Order: orderRes, Err: err}, nil
	}
}

func makeChangeStatusEndpoint(s order.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangeStatusRequest)
		err := s.ChangeStatus(ctx, req.ID, req.Status)
		return ChangeStatusResponse{Err: err}, nil
	}
}
