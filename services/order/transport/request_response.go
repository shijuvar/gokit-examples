package transport

import (
	"github.com/shijuvar/gokit-examples/services/order"
)

// CreateRequest holds the request parameters for the Create method.
type CreateRequest struct {
	Order order.Order
}

// CreateResponse holds the response values for the Create method.
type CreateResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}

// GetByIDRequest holds the request parameters for the GetByID method.
type GetByIDRequest struct {
	ID string
}

// GetByIDResponse holds the response values for the GetByID method.
type GetByIDResponse struct {
	Order order.Order `json:"order"`
	Err   error       `json:"error,omitempty"`
}

// ChangeStatusRequest holds the request parameters for the ChangeStatus method.
type ChangeStatusRequest struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// ChangeStatusResponse holds the response values for the ChangeStatus method.
type ChangeStatusResponse struct {
	Err error `json:"error,omitempty"`
}
