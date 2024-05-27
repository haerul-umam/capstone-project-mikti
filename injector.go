//go:build wireinject
//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/haerul-umam/tugas-mikti-kel1/app"
	"github.com/labstack/echo/v4"
)

func StartServer() *echo.Echo {
	wire.Build(
		app.Router,
	)
	return nil
}