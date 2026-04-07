# wedding-pocketbase

PocketBase backend for the wedding app suite. Runs as a single Go binary with an embedded SQLite database.

## Running

```bash
go run main.go serve        # start on :8090
go run main.go serve --http 0.0.0.0:8090  # bind to all interfaces
```

Admin UI is available at `http://127.0.0.1:8090/_/`.

## Collections

| Collection | Type | Description |
|---|---|---|
| `events` | base | Wedding events, each owned by a `users` record |
| `rsvps` | base | Guest RSVPs linked to an event, with check-in fields |
| `users` | auth | App users (event owners) |
| `_superusers` | auth | PocketBase admin accounts |

### Key fields

**events**: `name`, `slug` (unique), `date`, `owner` (relation to users), `check_in_secret` (auto-generated 32-char token, unique)

**rsvps**: `event` (relation), `name`, `email`, `attendance`, `guest_count`, `message`, `check_in_code` (auto-generated 12-char, unique), `checked_in`, `checked_in_at`

## Custom API

Two custom routes for the check-in app. These do **not** require PocketBase authentication -- the event's `check_in_secret` acts as a bearer secret.

### `POST /api/checkin/verify`

Validates a greeter's credentials and returns event info.

**Request**

```json
{
  "secret": "abc123...",
  "event_slug": "thomas-wedding"
}
```

**Response `200`**

```json
{
  "event": {
    "id": "abc123def456g",
    "name": "Thomas & Sarah's Wedding",
    "slug": "thomas-wedding",
    "date": "2026-06-15 00:00:00.000Z"
  }
}
```

**Response `401`** -- invalid secret or slug

```json
{ "message": "Invalid credentials" }
```

---

### `POST /api/checkin/scan`

Looks up a guest by their check-in QR code and marks them as checked in.

**Request**

```json
{
  "secret": "abc123...",
  "event_slug": "thomas-wedding",
  "check_in_code": "a1b2c3d4e5f6"
}
```

**Response `200`** -- guest checked in (or was already checked in)

```json
{
  "rsvp": {
    "name": "Jane Doe",
    "guest_count": 2,
    "checked_in": true,
    "checked_in_at": "2026-06-15T14:30:00Z",
    "already_checked_in": false
  }
}
```

`already_checked_in` is `true` if the guest was already checked in before this request.

**Response `401`** -- invalid secret or slug

```json
{ "message": "Invalid credentials" }
```

**Response `404`** -- no RSVP matches the check-in code for this event

```json
{ "message": "Guest not found" }
```

## Migrations

Migrations live in `pb_migrations/` and run automatically on startup:

| File | Description |
|---|---|
| `1744052400_init_collections.go` | Initial collections (users, rsvps) |
| `1744052401_add_checkin_fields.go` | Add check-in fields to rsvps |
| `1744052402_add_events_multi_tenancy.go` | Add events collection, link rsvps to events |
| `1744052403_add_checkin_secret.go` | Add `check_in_secret` field to events |
