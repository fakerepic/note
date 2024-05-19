package migrations

import (
	"fmt"
	"pbapp/config"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	conf := config.Load()
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)

		settings, _ := dao.FindSettings()
		settings.Meta.AppName = "pocketbase-notebook"
		settings.Meta.AppUrl = conf.APP_PUBLIC_URL
		settings.Logs.MaxDays = 2
		settings.Meta.SenderName = "Pocketbase Notebook"
		settings.Meta.SenderAddress = conf.SMTP_USER

		var port int
		fmt.Sscanf(conf.SMTP_PORT, "%d", &port)
		settings.Smtp.Port = port
		settings.Smtp.Host = conf.SMTP_HOST
		settings.Smtp.Username = conf.SMTP_USER
		settings.Smtp.Password = conf.SMTP_PASSWORD
		settings.Smtp.Tls = true
		settings.Smtp.Enabled = true

		return dao.SaveSettings(settings)
	}, nil)
}
