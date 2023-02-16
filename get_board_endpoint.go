package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type board struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Vendor string `json:"vendor"`
	Core   string `json:"core"`
}

func getBoard(w http.ResponseWriter, r *http.Request) {
	boardID := chi.URLParam(r, "boardID")

	var id int
	var name string
	var vendor string
	var core string
	err := db.QueryRow("select id,name,vendor,core from boards where id = $1", boardID).Scan(&id, &name, &vendor, &core)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	board := board{
		ID:     id,
		Name:   name,
		Vendor: vendor,
		Core:   core,
	}
	json.NewEncoder(w).Encode(board)
}
