package router

import (
	"github.com/GbSouza15/sistemaEstoque/pkg/utils"
	"github.com/gorilla/mux"
)

func WebhooksRoutes(r *mux.Router) error {

	r.HandleFunc("/webhook", utils.HandleWebhook)

	return nil
}
