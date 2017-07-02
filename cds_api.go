package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type Cluster struct {
	Name             string `json:"name"`
	Type             string `json:"type"`
	ConnectTimeoutMs int    `json:"connect_timeout_ms"`
	LBType           string `json:"lb_type"`
	ServiceName      string `json:"service_name"`
}
type Clusters struct {
	Clusters []Cluster `json:"clusters"`
}

// https://lyft.github.io/envoy/docs/configuration/cluster_manager/cds.html#get--v1-clusters-(string- service_cluster)-(string- service_node)
func getCluster(w http.ResponseWriter, r *http.Request) {
	_ = chi.URLParam(r, "service_cluster")
	_ = chi.URLParam(r, "service_node")

	result := Clusters{
		Clusters: []Cluster{
			Cluster{
				Name:             "plus",
				Type:             "sds",
				ConnectTimeoutMs: 250,
				LBType:           "round_robin",
				ServiceName:      "plus",
			},
			Cluster{
				Name:             "zipkin",
				Type:             "sds",
				ConnectTimeoutMs: 250,
				LBType:           "round_robin",
				ServiceName:      "zipkin",
			},
		},
	}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
