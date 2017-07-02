package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type RouteConfig struct {
	VirtualHosts []VirtualHost `json:"virtual_hosts"`
}

type VirtualHost struct {
	Name    string   `json:"name"`
	Domains []string `json:"domains"`
	Routes  []Route  `json:"routes"`
}

type Route struct {
	Prefix  string `json:"prefix"`
	Cluster string `json:"cluster"`
}

// https://lyft.github.io/envoy/docs/configuration/http_conn_man/rds.html#get--v1-routes-(string- route_config_name)-(string- service_cluster)-(string- service_node)
func getRoutes(w http.ResponseWriter, r *http.Request) {
	_ = chi.URLParam(r, "route_config_name")
	_ = chi.URLParam(r, "service_cluster")
	_ = chi.URLParam(r, "service_node")

	result := RouteConfig{
		VirtualHosts: []VirtualHost{
			VirtualHost{
				Name:    "plus_svc",
				Domains: []string{"*"},
				Routes: []Route{
					Route{Prefix: "/", Cluster: "plus"},
				},
			},
		},
	}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
