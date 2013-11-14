package main

import (
	"flag"
	"log"
	"net/http"
)

var bind = flag.String("bind", ":http", "address/port to bind to")
var message = flag.String("message", "Service unavailable.", "message text")
var status = flag.Int("status", http.StatusServiceUnavailable, "status code")

func main() {
	flag.Parse()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, *message, *status)
	})
	err := http.ListenAndServe(*bind, handler)
	if err != nil {
		log.Fatal(err)
	}
}
