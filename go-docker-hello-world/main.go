package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	handler "./handler"
)

func main() {

	var port string
	flag.StringVar(&port, "port", "8080", "Port to listen on")
	flag.Parse()

	r := handler.CreateRoutes()

	fmt.Printf("Server started on port %v...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
	fmt.Printf("Done!\n")
}
