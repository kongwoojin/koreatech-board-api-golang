package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "koreatech-board-api/docs"
	"koreatech-board-api/routes"
)

//	@title			KOREATECH board REST API
//	@version		1.0
//	@description	This is unofficial version of KOREATECH board REST API

//	@contact.name	Developer
//	@contact.email	kongwoojin03@gmail.com

//	@BasePath	/v3
func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.APIRouter(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
