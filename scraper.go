package scraper

// Go web scraper

import (
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	log.Println("Ping")
	w.Write([]byte("ping"))
}

func main() {
	addr := ":7171"
	http.HandleFunc("/ping", ping)

	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
