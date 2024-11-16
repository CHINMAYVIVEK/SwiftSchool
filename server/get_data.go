package server

import (
	"encoding/json"
	"net/http"

	"github.com/chinmayvivek/SwiftSchool/helper"
)

// Sample data structure
type Data struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}

func getData(w http.ResponseWriter, r *http.Request) {
	// Sample data to return
	data := []Data{
		{ID: 1, Value: "Item 1"},
		{ID: 2, Value: "Item 2"},
		{ID: 3, Value: "Item 3"},
	}
	helper.SugarObj.Info("data", data)

	// Set the response header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON data to the response
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
