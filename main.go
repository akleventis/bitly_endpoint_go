package main

import (
	"bitly_server_go/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	handler := handlers.New()
	router := mux.NewRouter()

	router.HandleFunc("/clicks", handler.GetClicks).Methods("GET")
	http.ListenAndServe(":8000", router)
}
