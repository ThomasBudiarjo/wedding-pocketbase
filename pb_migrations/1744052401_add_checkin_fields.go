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
        "id": "pbc_1102297273",
        "listRule": null,
        "viewRule": null,
        "createRule": "",
        "updateRule": null,
        "deleteRule": null,
        "name": "rsvps",
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
                "id": "text1579384326",
                "max": 0,
                "min": 0,
                "name": "name",
                "pattern": "",
                "presentable": false,
                "primaryKey": false,
                "required": true,
                "system": false,
                "type": "text"
            },
            {
                "exceptDomains": null,
                "hidden": false,
                "id": "email3885137012",
                "name": "email",
                "onlyDomains": null,
                "presentable": false,
                "required": false,
                "system": false,
                "type": "email"
            },
            {
                "hidden": false,
                "id": "select1843596689",
                "maxSelect": 1,
                "name": "attendance",
                "presentable": false,
                "required": true,
                "system": false,
                "type": "select",
                "values": [
                    "attending",
                    "not_attending"
                ]
            },
            {
                "hidden": false,
                "id": "number2418529142",
                "max": null,
                "min": null,
                "name": "guest_count",
                "onlyInt": false,
                "presentable": false,
                "required": false,
                "system": false,
                "type": "number"
            },
            {
                "convertURLs": false,
                "hidden": false,
                "id": "editor3065852031",
                "maxSize": 0,
                "name": "message",
                "presentable": false,
                "required": false,
                "system": false,
                "type": "editor"
            },
            {
                "autogeneratePattern": "[a-z0-9]{12}",
                "hidden": false,
                "id": "text_checkin_code",
                "max": 12,
                "min": 12,
                "name": "check_in_code",
                "pattern": "^[a-z0-9]+$",
                "presentable": false,
                "primaryKey": false,
                "required": true,
                "system": false,
                "type": "text"
            },
            {
                "hidden": false,
                "id": "bool_checked_in",
                "name": "checked_in",
                "presentable": false,
                "required": false,
                "system": false,
                "type": "bool"
            },
            {
                "hidden": false,
                "id": "date_checked_in_at",
                "max": "",
                "min": "",
                "name": "checked_in_at",
                "presentable": false,
                "required": false,
                "system": false,
                "type": "date"
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
            "CREATE UNIQUE INDEX ` + "`" + `idx_rsvps_check_in_code` + "`" + ` ON ` + "`" + `rsvps` + "`" + ` (` + "`" + `check_in_code` + "`" + `)"
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
