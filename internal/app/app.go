package app

import (
	"fmt"
	"net/http"
)

type App struct {
	Server *http.Server
}

func (a *App) Start() {
	fmt.Println("Listening on port 8080")
	a.Server.ListenAndServe()
}

func NewApp(server *http.Server) *App {
	return &App{Server: server}
}
