package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccountPrivate struct {
	AccountId primitive.ObjectID `json:"_id" bson:"_id"`
	Password  string             `json:"password" bson:"password"`
}

const (
	AccountPrivateCollectionName = "account_private"
)
