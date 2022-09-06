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
	file.Path = data.Path
	file.CreatedAt = time.Now().UTC().UnixMicro()
	updatedAt := time.Now().UTC().UnixMicro()
	file.UpdatedAt = &updatedAt

	countFile, err := repo.Collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	doc := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "path", Value: data.Path},
			{Key: "ext", Value: file.Ext},
			{Key: "created_at", Value: file.CreatedAt},
		}},
	}

	if countFile > 0 {
		doc = bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "path", Value: data.Path},
				{Key: "ext", Value: file.Ext},
				{Key: "updated_at", Value: file.UpdatedAt},
			}},
		}
	}

	_, err = repo.Collection.UpdateOne(ctx, filter, doc, opts)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &file, http.StatusOK, nil
}
