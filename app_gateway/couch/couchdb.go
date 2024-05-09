package couch

import (
	"log"
	"pbapp/config"

	"github.com/fjl/go-couchdb"
)

func Init(c config.Config) *couchdb.Client {
	rawurl := c.CouchAdminUrl()
	couch, err := couchdb.NewClient(rawurl, nil)
	if err != nil {
		log.Fatal(err)
	}
	return couch
}
