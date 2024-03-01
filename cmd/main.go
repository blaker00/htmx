package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/blaker00/htmx/handler"
)

func main() {
	http.HandleFunc("/", handler.RootHandler)

	fmt.Println("Starting webserver on port 8080.")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
