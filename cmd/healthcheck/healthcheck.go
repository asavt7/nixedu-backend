package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	url := flag.String("url", "http://localhost:8080/health", "url to send GET request for healthcheck")
	got, err := http.Get(*url)
	if err != nil || got.StatusCode > 299 || got.StatusCode < 200 {
		log.Fatal(err)
	}
}
