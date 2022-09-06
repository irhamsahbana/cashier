package domain

import (
	"context"
	"mime/multipart"
)

type File struct {
	UUID         string       `bson:"uuid"`
	BranchUUID   string       `bson:"branch_uuid"`
	FileableUUID string       `bson:"fileable_uuid"`
	FileableType FileableType `bson:"fileable_type"`
	Ext          string       `bson:"ext"`
	Path         string       `bson:"path"`
	CreatedAt    int64        `bson:"created_at"`
	UpdatedAt    *int64       `bson:"updated_at,omitempty"`
	DeletedAt    *int64       `bson:"deleted_at,omitempty"`
}

type FileUsecaseContract interface {
	UploadFile(ctx context.Context, branchId string, file *multipart.FileHeader, req *UploadFileRequest) (*UploadFileResponse, int, error)
}

type FileRepositoryContract interface {
	UploadFile(ctx context.Context, branchId string, data *File) (*File, int, error)
}
