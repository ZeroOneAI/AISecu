package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type Account struct {
	Id           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RegistryType string             `json:"registry_type" bson:"registry_type"`
	RegistryUrl  string             `json:"registry_url" bson:"registry_url"`
	Username     string             `json:"username" bson:"username"`
	Nickname     string             `json:"nickname" bson:"nickname"`
}

const (
	AccountCollectionName = "account"
)

var (
	AccountDefaultIndex = mongo.IndexModel{
		Keys:    bsonx.Doc{{Key: "username", Value: bsonx.Int32(1)}, {Key: "registry_type", Value: bsonx.Int32(1)}, {Key: "registry_url", Value: bsonx.Int32(1)}},
		Options: options.Index().SetUnique(true),
	}
)
