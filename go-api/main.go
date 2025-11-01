package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	mux := newMux()
	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/items/", itemsHandler)
	return mux
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"message": "Hello, world"})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	// Expect path /items/{id}
	idStr := strings.TrimPrefix(r.URL.Path, "/items/")
	// if there's trailing slash, remove it
	idStr = strings.TrimSuffix(idStr, "/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid item id"})
		return
	}
	q := r.URL.Query().Get("q")
	resp := map[string]interface{}{"item_id": id}
	if q == "" {
		resp["q"] = nil
	} else {
		resp["q"] = q
	}
	writeJSON(w, http.StatusOK, resp)
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
