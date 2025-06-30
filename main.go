package main

import (
	"log"
	"net/http"
	"time"
)
func main(){
	urls:=[]string {
		"https://localhost:5000",
		"https://localhost:4000",
		"https://localhost:8000",
	}
	var backends []*Container
	for _,addr := range urls{
		backend,err := newContainer(addr)
		if err != nil{
			log.Fatalf("Invalid url"+addr)
		}
		backends = append(backends, backend)
	}
	checkHealthStatus(backends)
}