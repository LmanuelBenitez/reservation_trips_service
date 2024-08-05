package models

import (
	"github.com/labstack/echo/v4"
)

type Context struct {
	Context     echo.Context
	RequestData interface{}
}
