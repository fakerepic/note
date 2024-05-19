package main

import (
	"log"
	"os"
	"strings"

	// couch admin instance:
	"pbapp/config"
	"pbapp/couch"

	// to be bind to pocketbase event hooks:
	"pbapp/hooks/couchdb_peruser"
	"pbapp/hooks/proxy"
	"pbapp/hooks/proxy/mw"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"

	_ "pbapp/migrations"
)

func main() {
	conf := config.Load()

	app := pocketbase.New()

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

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
