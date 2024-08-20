package server

import (
	"fmt"
	"net/http"
	"template_soa/libs/logger"
	"template_soa/middleware"
	"template_soa/models"
	"template_soa/server/di"
	"template_soa/server/routes"
	"time"

	"github.com/labstack/echo/v4"
	middle "github.com/labstack/echo/v4/middleware"
	"go.uber.org/dig"
)

var server *echo.Echo

func InitServer() {
	server = echo.New()
	server.HideBanner = true

	invokeContainer()
}

func invokeContainer() {
	container := di.BuildContainer()

	setupErrorHandler()
	setupRoutes(container)
	err := container.Invoke(func(config *models.Configuration) {
		setupMiddleware(config)
		startServer(config.Server.Port)
	})

	if err != nil {
		logger.Fatal("Server", "InvokeContainer", err.Error())
	}
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

func setupRoutes(container *dig.Container) {
	routes.SetRoutes(container)
	for _, route := range routes.Routes {
		if route.MiddlewareFunc != nil {
			server.Add(route.Method, route.Pattern, route.HandlerFunc, route.MiddlewareFunc).Name = route.Name
		} else {
			server.Add(route.Method, route.Pattern, route.HandlerFunc).Name = route.Name
		}
	}
}

func setupMiddleware(configuration *models.Configuration) {
	server.Use(
		middle.CORSWithConfig(
			middle.CORSConfig{
				AllowOrigins: configuration.Server.Cors,
				AllowMethods: []string{http.MethodGet},
				AllowHeaders: []string{"X-Requested-With", "Content-type", "Authorization"},
			}),
		middleware.Logger)
}

func startServer(port int) {
	server.Logger.Fatal(server.Start(fmt.Sprintf("%d:", port)))
}
