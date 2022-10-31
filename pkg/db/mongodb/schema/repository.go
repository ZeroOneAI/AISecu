package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type Repository struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	AccountId primitive.ObjectID `json:"account_id" bson:"account_id"`
}

const (
	RepositoryCollectionName = "repository"
)

var (
	RepositoryDefaultIndex = mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "account_id", Value: bsonx.Int32(1)}, {Key: "name", Value: bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	}
)
