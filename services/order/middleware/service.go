package middleware

import "github.com/shijuvar/gokit-examples/services/order"

// Middleware describes a service middleware.
type Middleware func(service order.Service) order.Service
