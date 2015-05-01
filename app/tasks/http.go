package app

import (
	// "fmt"
	"encoding/json"
	"net/http"
)

var (
	tasks = []Task{
		Task{
			"Do this",
			false,
		},
		Task{
			"Did this",
			true,
		},
	}
)

func init() {
	http.HandleFunc("/api/tasks", getAllJSON)
}

func getAllJSON(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(&tasks)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
