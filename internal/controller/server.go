package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewServer() *http.Server {
	handler := mux.NewRouter()
	handler.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "Hello World"}`))
	})
	return &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
}
