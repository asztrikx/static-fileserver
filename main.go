package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var path = flag.String("path", "", "Path to be served")
var port = flag.Int("port", 5000, "Without :")
var url = flag.String("url", "/", "Url's folder part, e.g. 127.0.0.1:5000/path/here")

func main() {
	flag.Parse()

	handler := http.FileServer(http.Dir(*path))
	http.Handle(*url, handler)

	log.Printf("Listening on :%d...\n", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatal(err)
	}
}
