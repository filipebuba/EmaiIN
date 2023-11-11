package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

type Product struct {
	ID   int
	Name string
}

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{},
	}
	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}
	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaigns", endpoints.HandlerError(handler.CampaignGetById))

	http.ListenAndServe(":3000", r)
}

// oque é um middleware?
//é um interceptador de requisições que pode interromper a requisição ou não
//e pode fazer alguma coisa antes de passar para o proximo middleware
