package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"time"
)

type Image struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RepositoryId primitive.ObjectID `json:"repository_id" bson:"repository_id"`
	Tag          string             `json:"tag" bson:"tag"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

const (
	ImageCollectionName = "image"
)

var (
	ImageDefaultIndex = mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "repository_id", Value: bsonx.Int32(1)}, {Key: "tag", Value: bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	}
)
