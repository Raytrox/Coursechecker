package main

import "github.com/gorilla/mux"

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/monitoring-units", GetAllUnits).Methods("GET", "OPTIONS")
	router.HandleFunc("monitor", courseHandler).Methods("POST", "OPTIONS")

	return router
}
