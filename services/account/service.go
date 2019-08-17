package account

import "context"

// Service describes the Account service.
type Service interface {
	CreateCustomer(ctx context.Context, customer Customer) error
	//AddMoneyToWallet(ctx context.Context, customerID string, amount float64) error
	//GetWalletBalance (ctx context.Context, id string) (float64, error)
	//MakePayment(ctx context.Context, customerID string,  amount float64) error
}
