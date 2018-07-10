package main

import (
	"fmt"
	"net/http"

	"github.com/marekparafianowicz/go-server/pages"
	"github.com/marekparafianowicz/go-server/sites"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Printf("Serving on port 8080")
	http.HandleFunc("/", pages.Index)
	http.HandleFunc("/sites", sites.Index)
	http.HandleFunc("/sites/show", sites.Show)
	http.ListenAndServe(":8080", nil)
}
