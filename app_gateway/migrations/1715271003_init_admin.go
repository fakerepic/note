package migrations

import (
	"pbapp/config"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	conf := config.Load()
	m.Register(func(db dbx.Builder) error {
		// add up queries...
		dao := daos.New(db)

		admin := &models.Admin{}
		admin.Email = conf.ADMIN_EMAIL
		admin.SetPassword(conf.ADMIN_PASSWORD)

		return dao.SaveAdmin(admin)
	}, func(db dbx.Builder) error { // optional revert operation

		dao := daos.New(db)

		admin, _ := dao.FindAdminByEmail(conf.ADMIN_EMAIL)
		if admin != nil {
			return dao.DeleteAdmin(admin)
		}

		// already deleted
		return nil
	},
	)
}
