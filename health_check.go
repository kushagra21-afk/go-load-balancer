package main
import (
	"log"
	"net/http"
	"time"
)

func checkHealthStatus(backends []*Container, path string, duration time.Duration){
	ticker := time.NewTicker(duration)
	go func(){
		<- ticker.C
		for _,b := range backends{
			go statusCheck(b)
		}
	}()
}
func statusCheck(b *Container){
	response, err := http.Get(b.parsed_url.String())
	if err != nil || response.StatusCode >= 500{
		b.setAlive(false)
		log.Printf("❌ %s marked as down\n", b.parsed_url)
	} else {
		b.setAlive(true)
		log.Printf("✅ %s is alive\n", b.parsed_url)
	}
	if response != nil {
		response.Body.Close()
	}
}
