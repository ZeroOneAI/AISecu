package mongodb

import (
	"context"
	"github.com/ZeroOneAI/AISecu/pkg/db/mongodb/schema"
	"github.com/ZeroOneAI/AISecu/pkg/dependencymanager"
	"github.com/ZeroOneAI/AISecu/pkg/dependencymanager/general_dependencymanager"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Controller struct {
	account        *mongo.Collection
	accountPrivate *mongo.Collection
	component      *mongo.Collection
	repository     *mongo.Collection
	image          *mongo.Collection
	// TODO 해당 부분 제거하기
	dependency dependencymanager.Interface
}

func NewController(protocol, hostname, port string, databaseName string, authCredential *options.Credential, dependencyHostname string) (*Controller, error) {
	var cancelFn context.CancelFunc

	clientOption := options.Client().ApplyURI(protocol + "://" + hostname + ":" + port)
	if authCredential != nil {
		clientOption = clientOption.SetAuth(*authCredential)
	}
	ctx := context.Background()
	ctx, cancelFn = context.WithTimeout(ctx, 10*time.Second)
	defer cancelFn()
	mongoClient, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		return nil, err
	}
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	db := mongoClient.Database(databaseName)
	return &Controller{
		account:        db.Collection(schema.AccountCollectionName),
		accountPrivate: db.Collection(schema.AccountPrivateCollectionName),
		component:      db.Collection(schema.ComponentCollectionName),
		repository:     db.Collection(schema.RepositoryCollectionName),
		image:          db.Collection(schema.ImageCollectionName),
		// TODO 해당 로직 제거하기
		dependency: general_dependencymanager.New(dependencyHostname),
	}, nil
}

func (c *Controller) InitAllCollectionIndex() error {
	ctxBg := context.Background()
	ctxAc, cancelFnAc := context.WithTimeout(ctxBg, 1*time.Second)
	defer cancelFnAc()
	_, err := c.account.Indexes().CreateOne(ctxAc, schema.AccountDefaultIndex)
	if err != nil {
		return err
	}
	ctxRp, cancelFnRp := context.WithTimeout(ctxBg, 1*time.Second)
	defer cancelFnRp()
	_, err = c.repository.Indexes().CreateOne(ctxRp, schema.RepositoryDefaultIndex)
	ctxImg, cancelFnImg := context.WithTimeout(ctxBg, 1*time.Second)
	defer cancelFnImg()
	_, err = c.image.Indexes().CreateOne(ctxImg, schema.ImageDefaultIndex)
	return err
}
