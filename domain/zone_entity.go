package domain

import "context"

type Zone struct {
	UUID        string   `bson:"uuid"`
	BranchUUID  string   `bson:"branch_uuid"`
	Name        string   `bson:"name"`
	Description *string  `bson:"description"`
	SpaceGroups []string `bson:"space_groups"`
	CreatedAt   int64    `bson:"created_at"`
	UpdatedAt   *int64   `bson:"updated_at"`
}

type ZoneWithSpaceGroups struct {
	UUID        string       `bson:"uuid"`
	BranchUUID  string       `bson:"branch_uuid"`
	Name        string       `bson:"name"`
	Description *string      `bson:"description"`
	SpaceGroups []SpaceGroup `bson:"space_groups"`
	CreatedAt   int64        `bson:"created_at"`
	UpdatedAt   *int64       `bson:"updated_at"`
}

type ZoneUsecaseContract interface {
	UpsertZones(ctx context.Context, branchId string, req *ZoneUpsertRequest) (*ZonesResponse, int, error)
	Zones(ctx context.Context, branchId string) (*ZonesResponse, int, error)
}

type ZoneRepositoryContract interface {
	UpsertZones(ctx context.Context, branchId string, zones []Zone) ([]ZoneWithSpaceGroups, int, error)
	Zones(ctx context.Context, branchId string) ([]ZoneWithSpaceGroups, int, error)
}
