package pages

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	m := Message{"Hello in API"}
	b, err := json.Marshal(m)
	if err != nil {
		http.Error(w, "JSON serialization error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}
