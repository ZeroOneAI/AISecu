package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Component struct {
	Id             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	RepositoryId   primitive.ObjectID `json:"repository_id" bson:"repository_id"`
	IsAlwaysLatest bool               `json:"is_always_latest" bson:"is_always_latest"`
	ImageId        primitive.ObjectID `json:"image_id,omitempty" bson:"image_id,omitempty"`
}

const (
	// ComponentCollectionName 은 실제로 별도의 Collection 에 저장하지 않음.
	ComponentCollectionName = "component"
)
