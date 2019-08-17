package grpc

import (
	"context"
	"github.com/shijuvar/gokit-examples/services/account"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	oldcontext "golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/shijuvar/gokit-examples/services/account/transport"
	"github.com/shijuvar/gokit-examples/services/account/transport/pb"
)

// grpc transport service for Account service.
type grpcServer struct {
	createCustomer kitgrpc.Handler
	logger         log.Logger
}

// NewGRPCServer returns a new gRPC service for the provided Go kit endpoints
func NewGRPCServer(
	endpoints transport.Endpoints, options []kitgrpc.ServerOption,
	logger log.Logger,
) pb.AccountServer {
	errorLogger := kitgrpc.ServerErrorLogger(logger)
	options = append(options, errorLogger)

	return &grpcServer{
		createCustomer: kitgrpc.NewServer(
			endpoints.CreateCustomer, decodeCreateCustomerRequest, encodeCreateCustomerResponse, options...,
		),
		logger: logger,
	}
}

// Generate glues the gRPC method to the Go kit service method
func (s *grpcServer) CreateCustomer(ctx oldcontext.Context, req *pb.CreateCustomerRequest) (*pb.CreateCustomerResponse, error) {
	_, rep, err := s.createCustomer.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateCustomerResponse), nil
}

// decodeCreateCustomerRequest decodes the incoming grpc payload to our go kit payload
func decodeCreateCustomerRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateCustomerRequest)
	return transport.CreateCustomerRequest{
		Customer: account.Customer{
			Email:    req.Email,
			Password: req.Password,
			Phone:    req.Phone,
		},
	}, nil
}

// encodeCreateCustomerResponse encodes the outgoing go kit payload to the grpc payload
func encodeCreateCustomerResponse(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(transport.CreateCustomerResponse)
	err := getError(res.Err)
	if err == nil {
		return &pb.CreateCustomerResponse{}, nil
	}
	return nil, err
}

func getError(err error) error {
	switch err {
	case nil:
		return nil
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}
