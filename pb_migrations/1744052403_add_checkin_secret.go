package pb_migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		jsonData := `[
    {
        "id": "pbc_events",
        "listRule": "owner = @request.auth.id",
        "viewRule": "owner = @request.auth.id",
        "createRule": null,
        "updateRule": null,
        "deleteRule": null,
        "name": "events",
        "type": "base",
        "fields": [
            {
                "autogeneratePattern": "[a-z0-9]{15}",
                "hidden": false,
                "id": "text3208210256",
                "max": 15,
                "min": 15,
                "name": "id",
                "pattern": "^[a-z0-9]+$",
                "presentable": false,
                "primaryKey": true,
                "required": true,
                "system": true,
                "type": "text"
            },
            {
                "autogeneratePattern": "",
                "hidden": false,
                "id": "text_event_name",
                "max": 0,
                "min": 0,
                "name": "name",
                "pattern": "",
                "presentable": true,
                "primaryKey": false,
                "required": true,
                "system": false,
                "type": "text"
            },
            {
                "autogeneratePattern": "",
                "hidden": false,
                "id": "text_event_slug",
                "max": 0,
                "min": 0,
                "name": "slug",
                "pattern": "^[a-z0-9-]+$",
                "presentable": false,
                "primaryKey": false,
                "required": true,
                "system": false,
                "type": "text"
            },
            {
                "hidden": false,
                "id": "date_event_date",
                "max": "",
                "min": "",
                "name": "date",
                "presentable": false,
                "required": false,
                "system": false,
                "type": "date"
            },
            {
                "cascadeDelete": false,
                "collectionId": "_pb_users_auth_",
                "hidden": false,
                "id": "relation_event_owner",
                "maxSelect": 1,
                "minSelect": 0,
                "name": "owner",
                "presentable": false,
                "required": true,
                "system": false,
                "type": "relation"
            },
            {
                "autogeneratePattern": "[a-z0-9]{32}",
                "hidden": false,
                "id": "text_checkin_secret",
                "max": 32,
                "min": 32,
                "name": "check_in_secret",
                "pattern": "^[a-z0-9]+$",
                "presentable": false,
                "primaryKey": false,
                "required": true,
                "system": false,
                "type": "text"
            },
            {
                "hidden": false,
                "id": "autodate2990389176",
                "name": "created",
                "onCreate": true,
                "onUpdate": false,
                "presentable": false,
                "system": false,
                "type": "autodate"
            },
            {
                "hidden": false,
                "id": "autodate3332085495",
                "name": "updated",
                "onCreate": true,
                "onUpdate": true,
                "presentable": false,
                "system": false,
                "type": "autodate"
            }
        ],
        "indexes": [
            "CREATE UNIQUE INDEX ` + "`" + `idx_events_slug` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `slug` + "`" + `)",
            "CREATE UNIQUE INDEX ` + "`" + `idx_events_check_in_secret` + "`" + ` ON ` + "`" + `events` + "`" + ` (` + "`" + `check_in_secret` + "`" + `)"
        ],
        "system": false
    }
]`

		collections := []map[string]any{}
		if err := json.Unmarshal([]byte(jsonData), &collections); err != nil {
			return err
		}

		return app.ImportCollections(collections, false)
	}, func(app core.App) error {
		return nil
	})
}
