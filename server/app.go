package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yunussandikci/shareterm/server/handler"
	"github.com/yunussandikci/shareterm/server/utils"
)

type App struct {
	echo       *echo.Echo
	apiHandler *handler.ApiHandler
	webHandler *handler.WebHandler
}

func (a *App) Initialize() {
	a.webHandler = handler.NewWebHandler()
	a.apiHandler = handler.NewApiHandler()
	a.echo = echo.New()
	a.echo.Use(middleware.Logger())
	a.echo.Use(middleware.Recover())
	a.echo.Renderer = utils.NewRenderer("public")
	a.echo.GET("/web/:name", a.webHandler.Read)
	a.echo.POST("/api", a.apiHandler.Create)
}

func (a *App) Run() {
	utils.PrintStartInfo()
	a.echo.HideBanner = true
	a.echo.Logger.Fatal(a.echo.Start(fmt.Sprintf(":8080")))
}