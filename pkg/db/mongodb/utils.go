package mongodb

import (
	"context"
	"github.com/ZeroOneAI/AISecu/pkg/db/mongodb/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func deleteElemWithId(ctx context.Context, collection *mongo.Collection, id primitive.ObjectID) (bool, error) {
	filter := bson.M{"_id": id}
	delCnt, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	if delCnt == nil {
		return false, errors.InvalidReturnFromDB
	}
	if delCnt.DeletedCount == 0 {
		return true, errors.NoExist
	}
	return true, nil
}

func deleteElemByFilter(ctx context.Context, collection *mongo.Collection, filter interface{}) (bool, error) {
	delCnt, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return false, err
	}
	if delCnt == nil {
		return false, errors.InvalidReturnFromDB
	}
	return true, nil
}

func insertElement[T any](ctx context.Context, collection *mongo.Collection, element T) (primitive.ObjectID, error) {
	var objId primitive.ObjectID
	id, err := collection.InsertOne(ctx, element)
	if err != nil {
		return objId, err
	}
	var ok bool
	objId, ok = id.InsertedID.(primitive.ObjectID)
	if !ok {
		return objId, errors.InvalidReturnFromDB
	}
	return objId, nil
}

func listElementByFilter[T any](ctx context.Context, sp *int64, limit *int64, collection *mongo.Collection, filter interface{}) ([]T, error) {
	opts := options.Find()
	if sp != nil {
		opts.SetSkip(*sp)
	}
	if limit != nil {
		opts.SetLimit(*limit)
	}

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	elemList := make([]T, 0)
	err = cursor.All(ctx, &elemList)
	if err != nil {
		return nil, err
	}
	return elemList, nil
}

func getElementById[T any](ctx context.Context, collection *mongo.Collection, id primitive.ObjectID, result *T) error {
	if result == nil {
		return errors.NilDetect
	}
	filter := bson.M{"_id": id}
	res := collection.FindOne(ctx, filter)
	return res.Decode(result)
}
