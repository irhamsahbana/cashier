package validator

import (
	"context"
	"lucy/cashier/bootstrap"
	"time"
)

func Exists(collection string, filter interface{}) error {
	c := context.Background()
	ctx, cancel := context.WithTimeout(c, bootstrap.App.Config.GetDuration("context.timeout")*time.Second)
	defer cancel()

	DB := bootstrap.App.Mongo.Database(bootstrap.App.Config.GetString("mongo.name"))
	coll := DB.Collection(collection)

	result := coll.FindOne(ctx, filter)
	if result.Err() != nil {
		return result.Err()
	}

	return nil
}
