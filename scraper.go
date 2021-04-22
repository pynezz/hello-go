package main

//
//			 Go web scraper
//				 0.0.1
//

import (
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	log.Println("Ping")
	w.Write([]byte("ping"))
}

func scraper() {
	addr := ":7171"
	http.HandleFunc("/ping", ping)

	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
