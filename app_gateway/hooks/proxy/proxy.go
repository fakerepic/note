package proxy

import (
	"log"
	"net/url"
	"strings"

	"github.com/fatih/color"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase/core"
)

type Options struct {
	Skipper     middleware.Skipper
	AuthHandler echo.HandlerFunc
}

func Register(app core.App, raw_url string, options Options) error {
	url, err := url.Parse(raw_url)
	if err != nil {
		return err
	}

	skipper := options.Skipper
	if skipper == nil {
		skipper = middleware.DefaultSkipper
	}

	auth_handler := options.AuthHandler
	if auth_handler == nil {
		auth_handler = func(c echo.Context) error {
			return nil
		}
	}

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		enableProxy(e, url, skipper, auth_handler)
		return nil
	})

	return nil
}

func enableProxy(e *core.ServeEvent, url *url.URL, skipper middleware.Skipper, auth_handler echo.HandlerFunc) {
	e.Router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if skipper(c) {
				return next(c)
			}
			if err := auth_handler(c); err != nil {
				return err
			}
			return next(c)
		}
	})

	e.Router.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Skipper: skipper,
		Balancer: middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
			{
				URL: url,
			},
		}),
	}))

	date := new(strings.Builder)
	log.New(date, "", log.LstdFlags).Print()

	bold := color.New(color.Bold).Add(color.FgGreen)
	bold.Printf(
		"%s Setup proxy to %s\n",
		strings.TrimSpace(date.String()),
		color.CyanString("%s", url.String()),
	)
}
