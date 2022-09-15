package domain

import "time"

type ZonesResponse struct {
	Zones []ZoneResponse `json:"zones"`
}

type ZoneResponse struct {
	UUID        string               `json:"uuid"`
	BranchUUID  string               `json:"branch_uuid"`
	Name        string               `json:"name"`
	Description *string              `json:"description"`
	SpaceGroups []SpaceGroupResponse `json:"space_groups"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   *time.Time           `json:"updated_at"`
}

type ZoneUpsertRequest struct {
	Zones []ZoneRequest `json:"zones"`
}

type ZoneRequest struct {
	UUID        string   `json:"uuid"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	SpaceGroups []string `json:"space_groups"`
	CreatedAt   string   `json:"created_at"`
}
