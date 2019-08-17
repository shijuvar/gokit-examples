package transport

import "github.com/shijuvar/gokit-examples/services/account"

type (
	CreateCustomerRequest struct {
		Customer account.Customer
	}
	CreateCustomerResponse struct {
		Err error
	}
	AddMoneyToWalletRequest struct {
		CustomerID string
		Amount     float64
	}
	AddMoneyToWalletResponse struct {
		Err error
	}
	GetWalletBalanceRequest struct {
		CustomerID string
	}
	GetWalletBalanceResponse struct {
		Amount float64
		Err    error
	}
	MakePaymentRequest struct {
		CustomerID string
		Amount     float64
	}
	MakePaymentResponse struct {
		Err error
	}
)
