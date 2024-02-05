package router

import (
	"database/sql"

	repository "github.com/GbSouza15/sistemaEstoque/internal/app/repositories"
	"github.com/GbSouza15/sistemaEstoque/pkg/utils"
	"github.com/gorilla/mux"
)

func WebhooksRoutes(db *sql.DB, r *mux.Router) error {

	PaymentRepository := repository.NewPaymentRepository(db)
	wh := utils.NewWebhookHandle(PaymentRepository)

	r.HandleFunc("/webhook", wh.HandleWebhook)

	return nil
}
