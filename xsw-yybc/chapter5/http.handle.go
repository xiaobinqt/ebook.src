package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hellp, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
