package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	sm := http.NewServeMux()
	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received: ", r.Method, " ", r.URL.Path, " ")
		w.Header().Set("Content-Type", "application/json")

		if r.Header.Get("X-Delay-Server") == "golang" {
			time.Sleep(1 * time.Second)
		}

		w.Write([]byte(`{"message": "success go"}`))
	})

	sm.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received: ", r.Method, " ", r.URL.Path, " ")
		w.Header().Set("Content-Type", "application/json")
	})

	log.Println("Server started at host: 0.0.0.0:8090")
	log.Fatal(http.ListenAndServe("0.0.0.0:8090", sm))
}
