package mw

import (
	"strings"

	"github.com/labstack/echo/v5"
)

func DefaultSkipper(c echo.Context) bool {
	return strings.HasPrefix(c.Request().URL.Path, "/_") || strings.HasPrefix(c.Request().URL.Path, "/api")
}

func SkipperFactory(prefix []string) func(c echo.Context) bool {
	return func(c echo.Context) bool {
		for _, p := range prefix {
			if strings.HasPrefix(c.Request().URL.Path, p) {
				return true
			}
		}
		return false
	}
}
