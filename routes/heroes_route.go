package routes

import (
	"Chi-hero-API-with-GOPG/dto"
	"Chi-hero-API-with-GOPG/errors"
	"Chi-hero-API-with-GOPG/services"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-pg/pg/v10"
	"net/http"
)

func SetupHeroesRoute(r *chi.Mux) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Chi is up and running on pod xyz"))
	})

	r.Route("/heroes", func(r chi.Router) {
		r.Get("/", getHeroes)
		r.Get("/{heroID}", getHeroByID)
	})
}

func getHeroes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	pgdb, ok := r.Context().Value("DB").(*pg.DB)

	if !ok {
		errors.HandleDBFromContextErr(w)
		return
	}

	heroes, err := services.GetHeroes(pgdb)

	switch err.(type) {
	case nil:
		json.NewEncoder(w).Encode(heroes)
		w.WriteHeader(http.StatusOK)
	default:
		errors.HandleErr(w, err)
		return
	}
}

func getHeroByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	heroID := chi.URLParam(r, "heroID")

	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		errors.HandleDBFromContextErr(w)
		return
	}

	hero, err := services.GetHero(pgdb, heroID)
	if err != nil {
		errors.HandleErr(w, err)
		return
	}

	response := &dto.HeroResponse{
		Success: true,
		Error:   "",
		Hero:    hero,
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}
