package handler

import (
	"net/http"
	config "ticketing/config"
	structure "ticketing/structure"

	"github.com/labstack/echo/v4"
)

func Login_info_handler(ctx echo.Context) error {
	var responses config.Response
	requestBody := new(structure.RequestLogin)
	
	// Bind
	if err := ctx.Bind(requestBody); err != nil {
		responses.Message = "Bind Failed"
		return responses.ResponseMiddleware(ctx, http.StatusBadRequest)
	}

	// Validate
	if err := ctx.Validate(requestBody); err != nil {
		responses.Message = "Bad Request"
		responses.Error = err.Error()
		return responses.ResponseMiddleware(ctx, http.StatusBadRequest)
	}
	
	responses.Message = "Success"
	return responses.ResponseMiddleware(ctx, http.StatusOK)
}