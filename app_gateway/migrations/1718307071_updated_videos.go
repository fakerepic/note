package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/tools/types"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("5oy47pwapvyamqt")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("@request.auth.id = owner.id")

		collection.ViewRule = types.Pointer("@request.auth.id = owner.id")

		collection.CreateRule = types.Pointer("@request.auth.id = owner.id")

		collection.UpdateRule = types.Pointer("@request.auth.id = owner.id")

		collection.DeleteRule = types.Pointer("@request.auth.id = owner.id")

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("5oy47pwapvyamqt")
		if err != nil {
			return err
		}

		collection.ListRule = types.Pointer("")

		collection.ViewRule = types.Pointer("")

		collection.CreateRule = types.Pointer("")

		collection.UpdateRule = types.Pointer("")

		collection.DeleteRule = types.Pointer("")

		return dao.SaveCollection(collection)
	})
}
