package models

import "github.com/labstack/echo/v4"

type Route struct {
	Method         string
	Name           string
	Pattern        string
	HandlerFunc    echo.HandlerFunc
	MiddlewareFunc echo.HandlerFunc
}
