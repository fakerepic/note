package couchdb_peruser

import (
	"fmt"
	"log"
	"slices"

	"github.com/fjl/go-couchdb"
	"github.com/pocketbase/pocketbase/core"
)

type User struct {
	Id       string `db:"id" json:"id"`
}

func Register(app core.App, couch *couchdb.Client) error {
	app.OnRecordAfterConfirmVerificationRequest().Add(
		func(e *core.RecordConfirmVerificationEvent) error {
			user_id := e.Record.Id
			CreateUserDBForSync(couch, user_id)
			return nil
		},
	)

	// TODO: ensure userdb for all verified users
	app.OnBeforeServe().Add(func(_ *core.ServeEvent) error {
		user := []User{}
		app.Dao().DB().NewQuery("SELECT id FROM user WHERE verified=true").All(&user)
		for _, u := range user {
			CreateUserDBForSync(couch, u.Id)
		}
		return nil
	})
	return nil
}

func CreateUserDBForSync(couch *couchdb.Client, userid string) {
	db_name := fmt.Sprintf("userdb-%s", userid)

	db, err := couch.EnsureDB(db_name)
	if err != nil {
		log.Fatal(err)
	}

	secobj, _ := db.Security()

	if !slices.Contains(secobj.Admins.Names, userid) {
		secobj.Admins.Names = append(secobj.Admins.Names, userid)
	}
	if !slices.Contains(secobj.Members.Names, userid) {
		secobj.Members.Names = append(secobj.Members.Names, userid)
	}

	db.PutSecurity(secobj)
}
