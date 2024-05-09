package main

import (
	"log"

	// couch admin instance:
	"pbapp/config"
	"pbapp/couch"

	// to be bind to pocketbase event hooks:
	"pbapp/hooks/couchdb_peruser"
	"pbapp/hooks/proxy"
	"pbapp/hooks/proxy/mw"

	"github.com/pocketbase/pocketbase"
)

func main() {
	conf := config.Load()

	app := pocketbase.New()

	couch := couch.Init(conf)

	// register couchdb per-user action
	couchdb_peruser.Register(app, couch)

	// register couchdb proxy authentication, expcept for /api, /_, /ai
	proxy.Register(app, conf.CouchUrl(),
		proxy.Options{
			Skipper:     mw.SkipperFactory([]string{"/api", "/_", "/ai"}),
			AuthHandler: mw.CouchAuthHandler,
		},
	)

	// register ai service proxy authentication, expcept for /api, /_
	proxy.Register(app, conf.AIServiceUrl(),
		proxy.Options{
			Skipper:     mw.SkipperFactory([]string{"/api", "/_"}),
			AuthHandler: mw.AIServiceAuthHandler,
		},
	)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
