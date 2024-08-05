package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"template_soa/models"

	"github.com/labstack/echo/v4"
)

func CheckTripRequest() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			request, err := parseBody(ctx)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, models.Response{
					Message:    "Bad Request",
					Details:    err.Error(),
					Status:     http.StatusBadRequest,
					ApiVersion: "1.0",
				})
			}

			ctx.Set("body", request)

			// context := models.Context{
			// 	Context:     ctx,
			// 	RequestData: request,
			// }

			return next(ctx)
		}
	}
}

func parseBody(ctx echo.Context) (*models.RequestTrip, error) {
	request := new(models.RequestTrip)

	defer ctx.Request().Body.Close()

	err := json.NewDecoder(ctx.Request().Body).Decode(&request)
	if err != nil {
		errorString := fmt.Sprintf("Error decoding request to RequestTrip model, with error: %v", err.Error())
		return nil, errors.New(errorString)
	}

	return request, nil
}
