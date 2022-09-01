package domain

import (
	"context"
	"errors"
)

type SpaceGroupShape string

const (
	SpaceGroupShape_CIRCLE SpaceGroupShape = "circle"
	SpaceGroupShape_SQUARE SpaceGroupShape = "square"
)

var ErrSpaceGroupShapeInvalid = errors.New("space group shape invalid")

type SpaceGroup struct {
	UUID       string          `bson:"uuid"`
	BranchUUID string          `bson:"branch_uuid"`
	Spaces     []Space         `bson:"space,omitempty"`
	Code       string          `bson:"code"`
	Shape      SpaceGroupShape `bson:"shape"`
	Pax        uint            `bson:"pax"`
	Floor      uint            `bson:"floor"`
	Smooking   bool            `bson:"smooking"`
	CreatedAt  int64           `bson:"created_at"`
	UpdatedAt  *int64          `bson:"updated_at,omitempty"`
	DeletedAt  *int64          `bson:"deleted_at,omitempty"`
}

type Space struct {
	UUID        string `bson:"uuid"`
	Number      int    `bson:"number"`
	Occupied    bool   `bson:"occupied"`
	Description string `bson:"description"`
	CreatedAt   int64  `bson:"created_at"`
	UpdatedAt   *int64 `bson:"updated_at,omitempty"`
	DeletedAt   *int64 `bson:"deleted_at,omitempty"`
}

type SpaceGroupUsecaseContract interface {
	UpsertSpaceGroup(ctx context.Context, branchId string, req *SpaceGroupUpsertRequest) (*SpaceGroupResponse, int, error)
	FindSpaceGroups(ctx context.Context, withTrashed bool) ([]SpaceGroupResponse, int, error)
	FindSpaceGroup(ctx context.Context, branchId, id string, withTrashed bool) (*SpaceGroupResponse, int, error)
	DeleteSpaceGroup(ctx context.Context, branchId, id string) (*SpaceGroupResponse, int, error)
}

type SpaceGroupRepositoryContract interface {
	UpsertSpaceGroup(ctx context.Context, data *SpaceGroup) (*SpaceGroup, int, error)
	// FindSpaceGroups(ctx context.Context, withTrashed bool) ([]SpaceGroup, int, error)
	FindSpaceGroup(ctx context.Context, branchId, id string, withTrashed bool) (*SpaceGroup, int, error)
	DeleteSpaceGroup(ctx context.Context, branchId, id string) (*SpaceGroup, int, error)
}
