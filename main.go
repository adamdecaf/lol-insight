package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/adamdecaf/lol-insights/routes"
)

const DefaultHttpPort = 8080

func main() {
	port := DefaultHttpPort

	// Setup and start the http server
	log.Printf("Starting http server on :%d\n", port)
	err := startHttpServer(port)
	if err != nil {
		header := fmt.Sprintf("error binding to port %d\n", port)
		log.Fatalf(header, err)
	}
}

func startHttpServer(port int) *error {
	http.Handle("/", http.FileServer(http.Dir("./html/")))
	http.HandleFunc("/ping", routes.Ping)

	listen := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(listen, nil)

	return &err
}
