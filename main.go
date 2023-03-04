package main

import (
	"Chi-hero-API-with-GOPG/db"
	"Chi-hero-API-with-GOPG/routes"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
	"log"
	"net/http"
	"os"
)

func main() {
	pgdb, err := StartDB()
	router := SetupRoutes(pgdb)
	port := FetchPortNumber()
	StartHTTPServer(err, port, router)
}

func StartDB() (*pg.DB, error) {
	pgdb, err := db.StartDB()
	if err != nil {
		log.Printf("error: %v", err)
		panic("error starting the database")

	}
	return pgdb, err
}

func SetupRoutes(pgdb *pg.DB) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))
	routes.SetupHeroesRoute(r)
	return r
}

func FetchPortNumber() string {
	port := os.Getenv("PORT")
	return port
}

func StartHTTPServer(err error, port string, router *chi.Mux) {
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("error from router %v\n", err)
		return
	}
}
