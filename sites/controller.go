package sites

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type sitesMessage struct {
	Sites []site
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	sites, err := allSites()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	sitesMessage := sitesMessage{sites}
	b, err := json.Marshal(sitesMessage)
	if err != nil {
		http.Error(w, "JSON serialization error", 500)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}
