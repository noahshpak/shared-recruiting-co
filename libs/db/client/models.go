// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package client

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type AuthUser struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

type UserEmailSyncHistory struct {
	UserID              uuid.UUID    `json:"user_id"`
	HistoryID           int64        `json:"history_id"`
	SyncedAt            time.Time    `json:"synced_at"`
	ExamplesCollectedAt sql.NullTime `json:"examples_collected_at"`
	CreatedAt           time.Time    `json:"created_at"`
	UpdatedAt           time.Time    `json:"updated_at"`
}

type UserOauthToken struct {
	UserID    uuid.UUID       `json:"user_id"`
	Provider  string          `json:"provider"`
	Token     json.RawMessage `json:"token"`
	IsValid   bool            `json:"is_valid"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type UserProfile struct {
	UserID    uuid.UUID `json:"user_id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Waitlist struct {
	UserID           uuid.UUID       `json:"user_id"`
	Email            string          `json:"email"`
	FirstName        string          `json:"first_name"`
	LastName         string          `json:"last_name"`
	LinkedinUrl      string          `json:"linkedin_url"`
	Responses        json.RawMessage `json:"responses"`
	CanCreateAccount bool            `json:"can_create_account"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}
