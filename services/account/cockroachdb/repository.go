package cockroachdb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"

	"github.com/shijuvar/gokit-examples/services/account"
)

var (
	ErrRepository = errors.New("unable to handle request")
)

type repository struct {
	db     *sql.DB
	logger log.Logger
}

// New returns a concrete repository backed by CockroachDB
func New(db *sql.DB, logger log.Logger) (account.Repository, error) {
	// return  repository
	return &repository{
		db:     db,
		logger: log.With(logger, "rep", "cockroachdb"),
	}, nil
}

// CreateOrder inserts a new order and its order items into db
func (repo *repository) CreateCustomer(ctx context.Context, customer account.Customer) error {

	// Insert order into the "orders" table.
	sql := `
			INSERT INTO customers (id, email, password, phone)
			VALUES ($1,$2,$3,$4)`
	_, err := repo.db.ExecContext(ctx, sql, customer.ID, customer.Email, customer.Password, customer.Phone)
	if err != nil {
		return err
	}
	return nil
}
