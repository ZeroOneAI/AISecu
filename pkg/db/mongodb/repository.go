package mongodb

import (
	"context"
	"github.com/ZeroOneAI/AISecu/pkg/db/mongodb/errors"
	"github.com/ZeroOneAI/AISecu/pkg/db/mongodb/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *Controller) AddRepository(ctx context.Context, accountId primitive.ObjectID, repositoryName string) (primitive.ObjectID, error) {
	newRepository := schema.Repository{
		Name:      repositoryName,
		AccountId: accountId,
	}

	var repositoryId primitive.ObjectID
	var ok bool
	id, err := c.repository.InsertOne(ctx, newRepository)
	if err != nil {
		return repositoryId, err
	}
	repositoryId, ok = id.InsertedID.(primitive.ObjectID)
	if !ok {
		return repositoryId, errors.InvalidReturnFromDB
	}
	return repositoryId, nil
}

func (c *Controller) GetRepository(ctx context.Context, repositoryId primitive.ObjectID) (*schema.Repository, error) {
	repository := &schema.Repository{}
	err := getElementById(ctx, c.repository, repositoryId, repository)
	return repository, err
}

func (c *Controller) DeleteRepository(ctx context.Context, repositoryId primitive.ObjectID) (bool, error) {
	imageDeleteFilter := bson.M{"repository_id": repositoryId}
	_, err := deleteElemByFilter(ctx, c.image, imageDeleteFilter)
	if err != nil {
		return false, err
	}
	// return deleteElemWithId(ctx, c.repository, repositoryId)
	notExist, err := deleteElemWithId(ctx, c.repository, repositoryId)
	if !notExist {
		return false, err
	}
	// TODO 해당 로직 제거하기
	err = c.dependency.DeleteMetricsRelatedToRepository(repositoryId.Hex())
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *Controller) ListRepository(ctx context.Context, sp *int64, limit *int64) ([]*schema.Repository, error) {
	return listElementByFilter[*schema.Repository](ctx, sp, limit, c.repository, bson.M{})
}

func (c *Controller) ListRepositoryByAccount(ctx context.Context, accountId primitive.ObjectID, sp *int64, limit *int64) ([]*schema.Repository, error) {
	filter := bson.M{"account_id": accountId}
	return listElementByFilter[*schema.Repository](ctx, sp, limit, c.repository, filter)
}
