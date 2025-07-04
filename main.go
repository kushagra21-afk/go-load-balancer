package main

import (
	"log"
	"net/http"
)

func main() {
	cfg, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	var backends []*Container
	for _, addr := range cfg.Backends {
		backend, err := newContainer(addr)
		if err != nil {
			log.Fatalf("Invalid url" + addr)
		}
		backends = append(backends, backend)
	}
	checkHealthStatus(backends, cfg.HealthCheck.Path, cfg.HealthCheck.IntervalDuration())
	lb := NewLoadBalancer(backends)
	log.Println("🚀 Load balancer running at", cfg.Listen)
	log.Fatal(http.ListenAndServe(cfg.Listen, lb))
}
