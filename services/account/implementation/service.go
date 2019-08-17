package implementation

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"

	"github.com/shijuvar/gokit-examples/services/account"
)

// service implements the Order Service
type service struct {
	repository account.Repository
	logger     log.Logger
}

// NewService creates and returns a new Account service instance
func NewService(rep account.Repository, logger log.Logger) account.Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

// Create makes an order
func (s *service) CreateCustomer(ctx context.Context, customer account.Customer) error {
	logger := log.With(s.logger, "method", "CreateCustomer")
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	customer.ID = id

	if err := s.repository.CreateCustomer(ctx, customer); err != nil {
		level.Error(logger).Log("err", err)
	}
	return nil
}
