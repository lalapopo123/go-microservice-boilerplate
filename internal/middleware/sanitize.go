package middleware

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/aditwar-man/go-microservice-boilerplate/pkg/sanitize"
)

// Sanitize and read request body to ctx for next use in easy json
func (mw *MiddlewareManager) Sanitize(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		body, err := io.ReadAll(ctx.Request().Body)
		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}
		defer ctx.Request().Body.Close()

		sanBody, err := sanitize.SanitizeJSON(body)
		if err != nil {
			return ctx.NoContent(http.StatusBadRequest)
		}

		ctx.Set("body", sanBody)
		return next(ctx)
	}
}
