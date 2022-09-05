package mongo

import (
	"context"
	"lucy/cashier/domain"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *fileMongoRepository) UploadFile(ctx context.Context, branchId string, data *domain.File) (*domain.File, int, error) {
	var file domain.File

	filter := bson.M{
		"uuid":          data.UUID,
		"branch_uuid":   branchId,
		"fileable_uuid": data.FileableUUID,
		"fileable_type": data.FileableType,
	}
	opts := options.Update().SetUpsert(true)

	file.UUID = data.UUID
	file.BranchUUID = branchId
	file.FileableUUID = data.FileableUUID
	file.FileableType = data.FileableType
	file.Ext = data.Ext
	file.CreatedAt = time.Now().UTC().UnixMicro()

	doc := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "uuid", Value: file.UUID},
			{Key: "branch_uuid", Value: file.BranchUUID},
			{Key: "fileable_uuid", Value: file.FileableUUID},
			{Key: "fileable_type", Value: file.FileableType},
			{Key: "ext", Value: file.Ext},
			{Key: "created_at", Value: file.CreatedAt},
		}},
	}

	_, err := repo.Collection.UpdateOne(ctx, filter, doc, opts)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &file, http.StatusOK, nil
}
