package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "koreatech-board-api/cmd/docs"
	"koreatech-board-api/cmd/routes"
	"os"
)

//	@title			KOREATECH board REST API
//	@version		1.0
//	@description	This is unofficial version of KOREATECH board REST API

//	@contact.name	Developer
//	@contact.email	kongwoojin03@gmail.com

// @BasePath	/v3
func main() {
	// Echo instance
	e := echo.New()

	f, err := os.Create("access.log")

	if err != nil {
		panic(err)
	}

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}: ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, user_agent=${user_agent}\n",
		Output: f,
	}))
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	routes.APIRouter(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
