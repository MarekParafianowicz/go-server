package sites

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type sitesMessage struct {
	Sites []site `json:"sites_attr"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	sites, err := allSites()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	sitesMessage := sitesMessage{sites}
	json, err := json.Marshal(sitesMessage)
	if err != nil {
		http.Error(w, "JSON serialization error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func Show(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	site, err := findSite(id)

	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(site)
	if err != nil {
		http.Error(w, "JSON serialization error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)

	atr := make(map[string]string)
	err := decoder.Decode(&atr)
	if err != nil {
		panic(err)
	}

	site, err := createSite(atr)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(site)
	if err != nil {
		http.Error(w, "JSON serialization error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" && r.Method != "PUT" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)

	atr := make(map[string]string)
	err := decoder.Decode(&atr)
	if err != nil {
		panic(err)
	}

	site, err := updateSite(atr)

	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	json, err := json.Marshal(site)
	if err != nil {
		http.Error(w, "JSON serialization error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(json))
}
