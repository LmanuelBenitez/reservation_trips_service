package server

import (
	"net/http"
	"template_soa/libs/logger"
	"template_soa/server/di"
	"time"

	"github.com/labstack/echo/v4"
)

var server *echo.Echo

func InitServer() {
	server = echo.New()
	server.HideBanner = true
}

func invokeContainer() {
	container := di.BuildContainer()

}

func setupErrorHandler() {
	server.HTTPErrorHandler = func(err error, context echo.Context) {
		statusCode := http.StatusInternalServerError

		if handler, ok := err.(*echo.HTTPError); ok {
			statusCode = handler.Code
		}

		logger.Request(context.Request().Method, statusCode, context.Request().URL.RequestURI(), time.Now())

		server.DefaultHTTPErrorHandler(err, context)
	}
}

func setupRoutes() {

}
