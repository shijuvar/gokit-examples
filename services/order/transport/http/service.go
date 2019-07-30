package http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/shijuvar/gokit-examples/services/order"
	"net/http"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/shijuvar/gokit-examples/services/order/transport"
)

var (
	ErrBadRouting = errors.New("bad routing")
)

// NewService wires Go kit endpoints to the HTTP transport.
func NewService(
	svcEndpoints transport.Endpoints, options []kithttp.ServerOption, logger log.Logger,
) http.Handler {
	// set-up router and initialize http endpoints
	var (
		r            = mux.NewRouter()
		errorLogger  = kithttp.ServerErrorLogger(logger)
		errorEncoder = kithttp.ServerErrorEncoder(encodeErrorResponse)
	)
	options = append(options, errorLogger, errorEncoder)
	//options := []kithttp.ServerOption{
	//	kithttp.ServerErrorLogger(logger),
	//	kithttp.ServerErrorEncoder(encodeError),
	//}
	// HTTP Post - /orders
	r.Methods("POST").Path("/orders").Handler(kithttp.NewServer(
		svcEndpoints.Create,
		decodeCreateRequest,
		encodeResponse,
		options...,
	))

	// HTTP Post - /orders/{id}
	r.Methods("GET").Path("/orders/{id}").Handler(kithttp.NewServer(
		svcEndpoints.GetByID,
		decodeGetByIDRequest,
		encodeResponse,
		options...,
	))

	// HTTP Post - /orders/status
	r.Methods("POST").Path("/orders/status").Handler(kithttp.NewServer(
		svcEndpoints.ChangeStatus,
		decodeChangeStausRequest,
		encodeResponse,
		options...,
	))
	return r
}

func decodeCreateRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.CreateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Order); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetByIDRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return transport.GetByIDRequest{ID: id}, nil
}

func decodeChangeStausRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req transport.ChangeStatusRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeErrorResponse(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case order.ErrOrderNotFound:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
