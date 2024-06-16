package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "5oy47pwapvyamqt",
			"created": "2024-06-13 18:19:28.501Z",
			"updated": "2024-06-13 18:19:28.501Z",
			"name": "videos",
			"type": "base",
			"system": false,
			"schema": [
				{
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
				},
				{
					"system": false,
					"id": "ei5f7fr2",
					"name": "owner",
					"type": "relation",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": null
					}
				},
				{
					"system": false,
					"id": "8snhldgo",
					"name": "id_local",
					"type": "text",
					"required": false,
					"presentable": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null,
						"pattern": ""
					}
				}
			],
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_EVGqNzT` + "`" + ` ON ` + "`" + `videos` + "`" + ` (\n  ` + "`" + `id_local` + "`" + `,\n  ` + "`" + `owner` + "`" + `\n)"
			],
			"listRule": "",
			"viewRule": "",
			"createRule": "",
			"updateRule": "",
			"deleteRule": "",
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("5oy47pwapvyamqt")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}
