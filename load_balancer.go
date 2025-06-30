package main

import (
	"net/http"
	"sync/atomic"
)

type load_balancer struct {
	backends []*Container
	current uint64
}

func NewLoadBalancer(backends []*Container) *load_balancer{
	return &load_balancer{backends: backends}
}
func (lb *load_balancer) getNextBackend() *Container {
	total := len(lb.backends)
	if total == 0 {
		return nil
	}
	for i:=0; i<total; i++{
		
		index := atomic.AddUint64(&lb.current, 1)
		b:= lb.backends[int(index)%total]
		if b.isAlive(){
			return b
		}

	}
	return nil
}
func (lb *load_balancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend := lb.getNextBackend()
	if backend == nil {
		http.Error(w, "Service unavailable: no backends available", http.StatusServiceUnavailable)
		return
	}
	backend.ReverseProxy.ServeHTTP(w, r)
}