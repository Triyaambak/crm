package routes

import (
	config "github.com/triyaambak/CRM/config"
	controller "github.com/triyaambak/CRM/controllers"

	"github.com/go-chi/chi/v5"
)

func SetUpRoutes(mux *chi.Mux, db *config.DbConfig) {

	c := controller.Controller{}

	mux.Get("/", c.ServeHomePage())
	mux.Get("/leads", c.GetAllLeads(db))
	mux.Get("/leads/${id}", c.GetLead(db))
}
