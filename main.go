package main

import (
	"github.com/camerontrewbase/godog/internal/app"
	"github.com/camerontrewbase/godog/internal/controller"
)

func main() {
	app.NewApp(controller.NewServer()).Start()
}
