package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"encoding/json"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}
	fmt.Printf("Listening on :%s", port)
	http.ListenAndServe(":"+port, router())
}

func router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/v1/registration/{service_name}", getServiceRegistration)
	r.Get("/v1/clusters/{service_cluster}/{service_node}", getCluster)
	r.Get("/v1/routes/{route_config_name}/{service_cluster}/{service_node}", getRoutes)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		result := map[string]float64{
			"answer": 404,
		}
		fmt.Printf("notfound at=info %s\n", r.URL.String())
		w.Header().Add("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	return r
}
