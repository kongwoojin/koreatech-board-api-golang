package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"koreatech-board-api/routes"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	routes.APIRouter(e)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
