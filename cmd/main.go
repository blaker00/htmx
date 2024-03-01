package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/blaker00/htmx/handler"
)

func main() {
	http.HandleFunc("/", handler.RootHandler)
	http.HandleFunc("/htmx", handler.HTMX)
	http.HandleFunc("/indicator", handler.Indicator)

	http.HandleFunc("/spinner.svg", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../static/spinner.svg")
	})

	fmt.Println("Starting webserver on port 8080.")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
