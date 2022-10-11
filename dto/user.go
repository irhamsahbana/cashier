package dto

import (
	"time"
)

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserResponse struct {
	UUID         string          `json:"uuid"`
	BranchUUID   string          `json:"branch_uuid"`
	RoleUUID     *string         `json:"role_uuid"`
	Branch       *BranchResponse `json:"branch,omitempty"`
	Name         string          `json:"name"`
	Role         string          `json:"role"`
	Email        *string         `json:"email,omitempty"`
	Password     *string         `json:"password,omitempty"`
	Phone        *string         `json:"phone,omitempty"`
	WA           *string         `json:"wa,omitempty"`
	ProfileUrl   *string         `json:"profile_url,omitempty"`
	Token        *string         `json:"access_token,omitempty"`
	RefreshToken *string         `json:"refresh_token,omitempty"`
	CreatedAt    *time.Time      `json:"created_at,omitempty"`
	UpdatedAt    *time.Time      `json:"updated_at,omitempty"`
	DeletedAt    *time.Time      `json:"deleted_at,omitempty"`
}

type CustomerResponse struct {
	UUID       string     `json:"uuid"`
	BranchUUID string     `json:"branch_uuid"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Phone      string     `json:"phone"`
	Address    string     `json:"address"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

// Requests

type CustomerUpserRequest struct {
	UUID       string `json:"uuid"`
	BranchUUID string `json:"branch_uuid"`
	Name       string `json:"name"`
	Dob        string `json:"date_of_birth"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	CreatedAt  string `json:"created_at"`
}
