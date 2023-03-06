package main

import (
	handler "ticketing/handler"
	structure "ticketing/structure"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Validator = &structure.RequestValidator{Validator: validator.New()}
	
	// Endpoint
	app.GET("/test", handler.Login_info_handler);

	// Start Service
	app.Start(":4000");
}