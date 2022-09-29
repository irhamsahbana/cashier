package domain

import "time"

type SpaceGroupResponse struct {
	UUID       string          `json:"uuid"`
	BranchUUID string          `json:"branch_uuid"`
	Name       string          `json:"name"`
	Spaces     []SpaceResponse `json:"spaces"`
	Code       string          `json:"code"`
	Shape      SpaceGroupShape `json:"shape"`
	Length     uint8           `json:"length"`
	Pax        uint            `json:"pax"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  *time.Time      `json:"updated_at"`
	DeletedAt  *time.Time      `json:"deleted_at"`
}

type SpaceResponse struct {
	UUID        string     `json:"uuid"`
	Number      int        `json:"number"`
	Occupied    bool       `json:"occupied"`
	Description *string    `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

// requests

type SpaceGroupUpsertRequest struct {
	UUID      string          `json:"uuid"`
	Code      string          `json:"code"`
	Name      string          `json:"name"`
	Shape     SpaceGroupShape `json:"shape"`
	Pax       uint            `json:"pax"`
	Floor     uint            `json:"floor"`
	Smooking  bool            `json:"smooking"`
	CreatedAt string          `json:"created_at"`
}

type SpaceCreateRequest struct {
	UUID        string  `json:"uuid"`
	Number      int     `json:"number"`
	Occupied    bool    `json:"occupied"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
}

type SpaceUpdateRequest struct {
	Number      int     `json:"number"`
	Occupied    bool    `json:"occupied"`
	Description *string `json:"description"`
}
