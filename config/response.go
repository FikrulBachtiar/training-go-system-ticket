package config

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    *interface{}  `json:"data"`
	Error   string      `json:"error"`
}

func (res Response) ResponseMiddleware(next echo.Context, http_status int) error {
	response := make(map[string]interface{})
	if res.Error == "" {
		var response_data interface{};

		if res.Data == nil {
			response_data = make(map[string]interface{})
		} else {
			response_data = res.Data
		}

		response["code"] = http_status
		response["message"] = res.Message
		response["data"] = response_data
	} else {
		response["code"] = http_status
		response["message"] = res.Message
		response["error"] = res.Error
	}
	return next.JSON(http_status, response)
}
