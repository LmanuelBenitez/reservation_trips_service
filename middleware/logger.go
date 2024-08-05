package middleware

import (
	"template_soa/libs/logger"
	"time"

	"github.com/labstack/echo/v4"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) (e error) {
		e = next(ctx)
		if e != nil {
			return e
		}

		logger.Request(
			ctx.Request().Method,
			ctx.Response().Status,
			ctx.Request().RequestURI,
			time.Now(),
		)

		return
	}
}
