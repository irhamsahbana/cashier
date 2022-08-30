package domain

import "time"

type UserRoleResponse struct {
	UUID      string     `json:"uuid"`
	Name      string     `json:"name"`
	Power     int        `json:"power"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type UserRoleUpsertrequest struct {
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	Power     int    `json:"power"`
	CreatedAt string `json:"created_at"`
}
