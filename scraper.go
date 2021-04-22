package main

//
//
//			 Go web scraper
//				 0.0.1
//
//

// Usage: go run scraper.go
// curl -s 127.0.0.1:7171/search?url=<URL>

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gocolly/colly"
)

var startTime time.Time

func ping(w http.ResponseWriter, r *http.Request) {
	log.Println("Ping")
	w.Write([]byte("ping"))
}

func scraper() {

	addr := ":7171"

	// Endpoints
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/search", getData)

	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func getData(w http.ResponseWriter, r *http.Request) {

	startTime = time.Now()

	// Verify that the URL param exists
	URL := r.URL.Query().Get("url")
	if URL == "" {
		log.Println("Missing URL..")
		return
	}
	log.Println("visiting", URL)

	// Create a collector that will handle the data collection from HTML
	c := colly.NewCollector()
	var response []string

	// OnHTML function allows using a callback function
	// when a certain HTML-tag i reached, call the anonymous
	// function, and add the info to the slice 'response'
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" {
			response = append(response, link)
		}
	})

	// Visit URL
	c.Visit(URL)

	// Response slice to JSON
	b, err := json.Marshal(response)
	if err != nil {
		log.Println("failed to serialize response:", err)
		return
	}

	// Add a header and write the body to the endpoint
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)

	t := time.Now()
	elapsed := t.Sub(startTime)

	log.Printf("Scraped in %v", elapsed)

}
