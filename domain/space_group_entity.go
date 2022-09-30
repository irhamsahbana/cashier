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
	Spaces     []Space         `bson:"spaces"`
	Name       string          `bson:"name"`
	Code       string          `bson:"code"`
	Length     uint8           `bson:"length"`
	Shape      SpaceGroupShape `bson:"shape"`
	Pax        uint            `bson:"pax"`
	Reservable bool            `bson:"reservable"`
	Disabled   bool            `bson:"disabled"`
	CreatedAt  int64           `bson:"created_at"`
	UpdatedAt  *int64          `bson:"updated_at"`
	DeletedAt  *int64          `bson:"deleted_at"`
}

type Space struct {
	UUID        string  `bson:"uuid"`
	Number      int     `bson:"number"`
	Occupied    bool    `bson:"occupied"`
	Description *string `bson:"description"`
	CreatedAt   int64   `bson:"created_at"`
	UpdatedAt   *int64  `bson:"updated_at"`
	DeletedAt   *int64  `bson:"deleted_at"`
}

type SpaceGroupUsecaseContract interface {
	UpsertSpaceGroup(ctx context.Context, branchId string, req *SpaceGroupUpsertRequest) (*SpaceGroupResponse, int, error)
	FindSpaceGroups(ctx context.Context, branchId string, withTrashed bool) ([]SpaceGroupResponse, int, error)
	FindSpaceGroup(ctx context.Context, branchId, id string, withTrashed bool) (*SpaceGroupResponse, int, error)
	DeleteSpaceGroup(ctx context.Context, branchId, id string) (*SpaceGroupResponse, int, error)

	CreateSpace(ctx context.Context, branchId, spaceGroupId string, req *SpaceCreateRequest) (*SpaceResponse, int, error)
	FindSpace(ctx context.Context, branchId, id string, withTrashed bool) (*SpaceResponse, int, error)
	DeleteSpace(ctx context.Context, branchId, id string) (*SpaceResponse, int, error)
	UpdateSpace(ctx context.Context, branchId, id string, req *SpaceUpdateRequest) (*SpaceResponse, int, error)
}

type SpaceGroupRepositoryContract interface {
	UpsertSpaceGroup(ctx context.Context, data *SpaceGroup) (*SpaceGroup, int, error)
	FindSpaceGroups(ctx context.Context, branchId string, withTrashed bool) ([]SpaceGroup, int, error)
	FindSpaceGroup(ctx context.Context, branchId, id string, withTrashed bool) (*SpaceGroup, int, error)
	DeleteSpaceGroup(ctx context.Context, branchId, id string) (*SpaceGroup, int, error)

	InsertSpace(ctx context.Context, branchId, spaceGroupId string, data *Space) (*SpaceGroup, int, error)
	FindSpace(ctx context.Context, branchId, id string, withTrashed bool) (*SpaceGroup, int, error)
	DeleteSpace(ctx context.Context, branchId, id string) (*SpaceGroup, int, error)
	UpdateSpace(ctx context.Context, branchId, id string, data *Space) (*SpaceGroup, int, error)
}
