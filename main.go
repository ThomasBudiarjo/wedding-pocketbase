package main

import (
	"log"
	"net/http"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"

	_ "wedding-pocketbase/pb_migrations"
)

func main() {
	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.POST("/api/checkin/verify", func(e *core.RequestEvent) error {
			data := struct {
				Secret    string `json:"secret" form:"secret"`
				EventSlug string `json:"event_slug" form:"event_slug"`
			}{}
			if err := e.BindBody(&data); err != nil {
				return e.BadRequestError("Invalid request body", err)
			}
			if data.Secret == "" || data.EventSlug == "" {
				return e.BadRequestError("Missing secret or event_slug", nil)
			}

			records, err := e.App.FindRecordsByFilter(
				"events",
				"slug = {:slug} && check_in_secret = {:secret}",
				"",
				1,
				0,
				dbx.Params{"slug": data.EventSlug, "secret": data.Secret},
			)
			if err != nil || len(records) == 0 {
				return e.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Invalid credentials",
				})
			}

			event := records[0]
			return e.JSON(http.StatusOK, map[string]any{
				"event": map[string]any{
					"id":   event.Id,
					"name": event.GetString("name"),
					"slug": event.GetString("slug"),
					"date": event.GetString("date"),
				},
			})
		})

		se.Router.POST("/api/checkin/scan", func(e *core.RequestEvent) error {
			data := struct {
				Secret      string `json:"secret" form:"secret"`
				EventSlug   string `json:"event_slug" form:"event_slug"`
				CheckInCode string `json:"check_in_code" form:"check_in_code"`
			}{}
			if err := e.BindBody(&data); err != nil {
				return e.BadRequestError("Invalid request body", err)
			}
			if data.Secret == "" || data.EventSlug == "" || data.CheckInCode == "" {
				return e.BadRequestError("Missing required fields", nil)
			}

			events, err := e.App.FindRecordsByFilter(
				"events",
				"slug = {:slug} && check_in_secret = {:secret}",
				"",
				1,
				0,
				dbx.Params{"slug": data.EventSlug, "secret": data.Secret},
			)
			if err != nil || len(events) == 0 {
				return e.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Invalid credentials",
				})
			}

			event := events[0]

			rsvps, err := e.App.FindRecordsByFilter(
				"rsvps",
				"check_in_code = {:code} && event = {:eventId}",
				"",
				1,
				0,
				dbx.Params{"code": data.CheckInCode, "eventId": event.Id},
			)
			if err != nil || len(rsvps) == 0 {
				return e.JSON(http.StatusNotFound, map[string]string{
					"message": "Guest not found",
				})
			}

			rsvp := rsvps[0]
			alreadyCheckedIn := rsvp.GetBool("checked_in")

			if !alreadyCheckedIn {
				rsvp.Set("checked_in", true)
				rsvp.Set("checked_in_at", time.Now().UTC().Format(time.RFC3339))
				if err := e.App.Save(rsvp); err != nil {
					return e.InternalServerError("Failed to check in guest", err)
				}
			}

			return e.JSON(http.StatusOK, map[string]any{
				"rsvp": map[string]any{
					"name":               rsvp.GetString("name"),
					"guest_count":        rsvp.GetInt("guest_count"),
					"checked_in":         true,
					"checked_in_at":      rsvp.GetString("checked_in_at"),
					"already_checked_in": alreadyCheckedIn,
				},
			})
		})

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
