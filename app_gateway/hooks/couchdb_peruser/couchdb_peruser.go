package couchdb_peruser

import (
	"fmt"
	"log"
	"slices"

	"github.com/fjl/go-couchdb"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
)

func Register(app core.App, couch *couchdb.Client) error {
	app.OnRecordAfterConfirmVerificationRequest().Add(
		func(e *core.RecordConfirmVerificationEvent) error {
			user_id := e.Record.Id
			CreateUserDBForSync(couch, "userdb", user_id)
			CreateUserDBForSync(couch, "userattachment", user_id)
			return nil
		},
	)

	// TODO: put this in a migration
	// ensure userdb for all verified users
	app.OnAfterBootstrap().Add(func(_ *core.BootstrapEvent) error {
		records := []models.BaseModel{}
		app.Dao().DB().NewQuery("SELECT id FROM users WHERE verified=true").All(&records)
		for _, u := range records {
			CreateUserDBForSync(couch, "userdb", u.Id)
			CreateUserDBForSync(couch, "userattachment", u.Id)
		}
		return nil
	})
	return nil
}

func CreateUserDBForSync(couch *couchdb.Client, prefix string, userid string) {
	db_name := fmt.Sprintf("%s-%s", prefix, userid)

	db, err := couch.EnsureDB(db_name)
	if err != nil {
		log.Fatal(err)
	}

	secobj, _ := db.Security()

	if secobj == nil {
		return;
	}
	if !slices.Contains(secobj.Admins.Names, userid) {
		secobj.Admins.Names = append(secobj.Admins.Names, userid)
	}
	if !slices.Contains(secobj.Members.Names, userid) {
		secobj.Members.Names = append(secobj.Members.Names, userid)
	}

	db.PutSecurity(secobj)
}
