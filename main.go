package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://marekparafianowicz:password@localhost/go-server?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}

type message struct {
	Data string
}

type sitesMessage struct {
	Sites []site
}

type site struct {
	ID  int
	URL string
}

func main() {
	fmt.Printf("Serving on port 8080")
	http.HandleFunc("/", index)
	http.HandleFunc("/sites", sites)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	m := message{"Hello in API"}
	b, err := json.Marshal(m)
	if err != nil {
		http.Error(w, "JSON serialization error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}

func sites(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	rows, err := db.Query("SELECT * FROM sites")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		fmt.Println(err)
		return
	}
	defer rows.Close()

	sites := make([]site, 0)
	for rows.Next() {
		site := site{}
		err := rows.Scan(&site.ID, &site.URL)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			fmt.Println(err)
			return
		}
		sites = append(sites, site)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
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
