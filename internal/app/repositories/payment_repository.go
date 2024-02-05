package repository

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type PaymentRepository struct {
	Db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db}
}

func (p *PaymentRepository) CreatePayment(paymentIdStripe, customerIdStripe string, createdPayment time.Time) error {

	paymentId := uuid.New()

	_, err := p.Db.Exec("INSERT INTO payments_history (id, payment_id, customer_id, payment_date) VALUES (?, ?, ?, ?)", paymentId, paymentIdStripe, customerIdStripe, createdPayment)
	if err != nil {
		return err
	}

	return nil
}
