package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("5oy47pwapvyamqt")
		if err != nil {
			return err
		}

		// update
		edit_video := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ihwjfcas",
			"name": "video",
			"type": "file",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"mimeTypes": [
					"video/mp4",
					"video/x-ms-wmv",
					"video/quicktime",
					"video/3gpp"
				],
				"thumbs": [],
				"maxSelect": 1,
				"maxSize": 10000000,
				"protected": false
			}
		}`), edit_video); err != nil {
			return err
		}
		collection.Schema.AddField(edit_video)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("5oy47pwapvyamqt")
		if err != nil {
			return err
		}

		// update
		edit_video := &schema.SchemaField{}
		if err := json.Unmarshal([]byte(`{
			"system": false,
			"id": "ihwjfcas",
			"name": "video",
			"type": "file",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {
				"mimeTypes": [
					"video/mp4",
					"video/x-ms-wmv",
					"video/quicktime",
					"video/3gpp"
				],
				"thumbs": [],
				"maxSelect": 1,
				"maxSize": 5242880,
				"protected": false
			}
		}`), edit_video); err != nil {
			return err
		}
		collection.Schema.AddField(edit_video)

		return dao.SaveCollection(collection)
	})
}
