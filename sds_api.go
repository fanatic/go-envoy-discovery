package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type Host struct {
	IPAddress           string `json:"ip_address"`                      // Upstream host
	Port                int    `json:"port"`                            // Upstream port
	Zone                string `json:"az,omitempty"`                    // Upstream zone
	Canary              bool   `json:"canary,omitempty"`                // Canary status
	LoadBalancingWeight int    `json:"load_balancing_weight,omitempty"` // Weight [1-100]
}
type ServiceRegistration struct {
	Hosts []Host `json:"hosts"`
}

// https://lyft.github.io/envoy/docs/configuration/cluster_manager/sds_api.html#get--v1-registration-(string- service_name)
func getServiceRegistration(w http.ResponseWriter, r *http.Request) {
	serviceName := chi.URLParam(r, "service_name")

	var host Host
	if serviceName == "plus" {
		host = Host{IPAddress: "127.0.0.1", Port: 3000}
	} else if serviceName == "zipkin" {
		host = Host{IPAddress: "192.168.3.38", Port: 9411}
	}
	result := ServiceRegistration{Hosts: []Host{host}}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
