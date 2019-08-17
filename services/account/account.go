package account

import "context"

type Customer struct {
	ID, Email, Password, Phone string
}
type Repository interface {
	CreateCustomer(ctx context.Context, customer Customer) error
}
