package main

import (
	"net/http"
)

type ID struct {
	id int64
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
	}
	rows, err := db.Query("SELECT * FROM project")

}
