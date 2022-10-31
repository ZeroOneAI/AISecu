package mongodb

import (
	"context"
	"github.com/ZeroOneAI/AISecu/pkg/db/mongodb/errors"
	"github.com/ZeroOneAI/AISecu/pkg/db/mongodb/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (c *Controller) GetImage(ctx context.Context, imageId primitive.ObjectID) (*schema.Image, error) {
	image := &schema.Image{}
	err := getElementById(ctx, c.image, imageId, image)
	return image, err
}

func (c *Controller) CreateOrUpdateImage(ctx context.Context, repositoryId primitive.ObjectID, tag string) (primitive.ObjectID, error) {
	imageId, err := c.getImage(ctx, repositoryId, tag)
	if err != nil {
		if err != errors.NoExist {
			return primitive.ObjectID{}, nil
		} else {
			return c.createImage(ctx, repositoryId, tag)
		}
	}
	return imageId, c.updateImage(ctx, imageId)
}

func (c *Controller) ListImageByRepositoryId(ctx context.Context, repositoryId primitive.ObjectID, sp *int64, limit *int64) ([]*schema.Image, error) {
	filter := bson.M{"repository_id": repositoryId}
	return listElementByFilter[*schema.Image](ctx, sp, limit, c.image, filter)
}

func (c *Controller) GetLatestImageByRepositoryId(ctx context.Context, repositoryId primitive.ObjectID) (*schema.Image, error) {
	opts := options.Find().SetSort(bson.D{{Key: "updated_at", Value: -1}}).SetLimit(1)
	filter := bson.M{"repository_id": repositoryId}

	cursor, err := c.image.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	imageList := make([]*schema.Image, 0)
	err = cursor.All(ctx, &imageList)
	if err != nil {
		return nil, err
	}
	if len(imageList) != 1 {
		return nil, errors.NoExist
	}
	return imageList[0], nil
}

func (c *Controller) createImage(ctx context.Context, repositoryId primitive.ObjectID, tag string) (primitive.ObjectID, error) {
	newImage := schema.Image{
		RepositoryId: repositoryId,
		Tag:          tag,
		UpdatedAt:    time.Now(),
	}
	return insertElement(ctx, c.image, newImage)
}

func (c *Controller) updateImage(ctx context.Context, imageId primitive.ObjectID) error {
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "updated_at", Value: time.Now()}}}}
	_, err := c.image.UpdateByID(ctx, imageId, update)
	return err
}

func (c *Controller) getImage(ctx context.Context, repositoryId primitive.ObjectID, tag string) (primitive.ObjectID, error) {
	var imageId primitive.ObjectID
	filter := bson.M{"repository_id": repositoryId, "tag": tag}
	image, err := listElementByFilter[*schema.Image](ctx, nil, nil, c.image, filter)
	if err != nil {
		return imageId, err
	}
	if len(image) != 1 {
		return imageId, errors.NoExist
	}
	imageId = image[0].Id
	return imageId, nil
}
