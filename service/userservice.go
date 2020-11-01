package service

import (
	"context"
	"github.com/StrandingHeart/GrpcDemo/db/mongo"
	"github.com/StrandingHeart/GrpcDemo/grpc/user"
	"time"
)

func InsertUser(ctx context.Context, in *user.UserInsertRequest) error {
	db, err := mongo.ConnectToDB(`mongodb://127.0.0.1:27017`, `test`, 30*time.Second, 3)
	if err != nil {
		return err
	}
	_, err = db.Collection(`user`).InsertOne(ctx, in.Data)
	if err != nil {
		return err
	}
	return nil
}
