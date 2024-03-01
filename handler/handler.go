package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("../static/index.html")
	if err != nil {
		http.Error(w, "Error: Missing index.html file.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	_, err = w.Write(file)
	if err != nil {
		http.Error(w, "error processing webpage", http.StatusInternalServerError)
	}

}

func HTMX(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("../static/htmx.html")
	if err != nil {
		http.Error(w, "Error: Missing htmx.html file.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	_, err = w.Write(file)
	if err != nil {
		http.Error(w, "error processing webpage", http.StatusInternalServerError)
	}

}

func Indicator(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	time.Sleep(5 * time.Second)
	fmt.Fprintf(w, "<h2>5 seconds elapsed</h2>")
}
