package dto

import "time"

type WaiterUpsertrequest struct {
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type WaiterResponse struct {
	UUID       string     `json:"uuid"`
	BranchUUID string     `json:"branch_uuid"`
	Name       string     `json:"name"`
	LastActive *time.Time `json:"last_active,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}
