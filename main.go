package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Serving on port 8080")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

type Message struct {
	Data string
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	m := Message{"Hello in API"}
	response, err := marshalJSON(m)
	if err != nil {
		http.Error(w, "JSON serialization error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(response))
}

func marshalJSON(m Message) ([]byte, error) {
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	return b, nil
}
