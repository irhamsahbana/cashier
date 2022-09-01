package domain

import "time"

type SpaceGroupResponse struct {
	UUID       string          `json:"uuid"`
	BranchUUID string          `json:"branch_uuid"`
	Spaces     []SpaceResponse `json:"space,omitempty"`
	Code       string          `json:"code"`
	Shape      SpaceGroupShape `json:"shape"`
	Pax        uint            `json:"pax"`
	Floor      uint            `json:"floor"`
	Smooking   bool            `json:"smooking"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  *time.Time      `json:"updated_at,omitempty"`
	DeletedAt  *time.Time      `json:"deleted_at,omitempty"`
}

type SpaceResponse struct {
	UUID        string     `json:"uuid"`
	Number      int        `json:"number"`
	Occupied    bool       `json:"occupied"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

type SpaceGroupUpsertRequest struct {
	UUID      string          `json:"uuid"`
	Code      string          `json:"code"`
	Shape     SpaceGroupShape `json:"shape"`
	Pax       uint            `json:"pax"`
	Floor     uint            `json:"floor"`
	Smooking  bool            `json:"smooking"`
	CreatedAt string          `json:"created_at"`
}
