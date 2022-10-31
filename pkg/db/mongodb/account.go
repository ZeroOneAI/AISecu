package mongodb

import (
	"context"
	"github.com/ZeroOneAI/AISecu/pkg/db/mongodb/errors"
	"github.com/ZeroOneAI/AISecu/pkg/db/mongodb/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *Controller) AddAccount(ctx context.Context, registryType, registryUrl, accountUsername, accountNickname, accountPassword string) (primitive.ObjectID, error) {
	accountId, err := c.addAccount(ctx, registryType, registryUrl, accountUsername, accountNickname)
	if err != nil {
		return accountId, err
	}

	_, err = c.addAccountPrivate(ctx, accountId, accountPassword)
	return accountId, err
}

func (c *Controller) GetAccount(ctx context.Context, accountId primitive.ObjectID) (*schema.Account, error) {
	account := &schema.Account{}
	err := getElementById(ctx, c.account, accountId, account)
	return account, err
}

func (c *Controller) GetAccountPrivate(ctx context.Context, accountId primitive.ObjectID) (*schema.AccountPrivate, error) {
	account := &schema.AccountPrivate{}
	err := getElementById(ctx, c.accountPrivate, accountId, account)
	return account, err
}

/*
func (c *Controller) GetAccountByAccountData(ctx context.Context, registryType, registryUrl, accountUsername string) (*schema.Account, error) {
	filter := bson.M{"registry_type": registryType, "registry_url": registryUrl, "username": accountUsername}
	res := c.account.FindOne(ctx, filter)
	result := &schema.Account{}
	err := res.Decode(result)
	return result, err
}
*/

func (c *Controller) UpdateAccountNickname(ctx context.Context, accountId primitive.ObjectID, nickname string) error {
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "nickname", Value: nickname}}}}
	_, err := c.account.UpdateByID(ctx, accountId, update)
	return err
}

func (c *Controller) UpdateAccountPassword(ctx context.Context, accountId primitive.ObjectID, password string) error {
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "password", Value: password}}}}
	_, err := c.accountPrivate.UpdateByID(ctx, accountId, update)
	return err
}

func (c *Controller) DeleteAccount(ctx context.Context, accountId primitive.ObjectID) (bool, error) {
	repositories, err := c.ListRepositoryByAccount(ctx, accountId, nil, nil)
	if err != nil {
		return false, err
	}
	for _, repository := range repositories {
		noExist, _ := c.DeleteRepository(ctx, repository.Id)
		if !noExist {
			return false, errors.Undefined
		}
	}
	noExist, err := c.deleteAccount(ctx, accountId)
	if !noExist {
		if err != nil {
			return noExist, err
		}
		return noExist, errors.Undefined
	}
	noExist, err = c.deleteAccountPrivate(ctx, accountId)
	if !noExist {
		if err != nil {
			return noExist, err
		}
		return noExist, errors.Undefined
	}
	return true, nil
}

func (c *Controller) ListAccount(ctx context.Context, sp *int64, limit *int64) ([]*schema.Account, error) {
	return listElementByFilter[*schema.Account](ctx, sp, limit, c.account, bson.M{})
}

func (c *Controller) addAccount(ctx context.Context, registryType, registryUrl, accountUsername, accountNickName string) (primitive.ObjectID, error) {
	if accountNickName == "" {
		if registryUrl == "" && registryType == "dockerhub" {
			registryUrl = "docker.io"
		}
		if registryUrl != "" {
			accountNickName = registryUrl + "/"
		}
		accountNickName += accountUsername
	}
	newAccount := schema.Account{
		RegistryType: registryType,
		RegistryUrl:  registryUrl,
		Username:     accountUsername,
		Nickname:     accountNickName,
	}

	return insertElement(ctx, c.account, newAccount)
}

func (c *Controller) addAccountPrivate(ctx context.Context, accountId primitive.ObjectID, password string) (primitive.ObjectID, error) {
	newAccountPrivate := schema.AccountPrivate{
		AccountId: accountId,
		Password:  password,
	}

	return insertElement(ctx, c.accountPrivate, newAccountPrivate)
}

func (c *Controller) deleteAccount(ctx context.Context, accountId primitive.ObjectID) (bool, error) {
	return deleteElemWithId(ctx, c.account, accountId)
}

func (c *Controller) deleteAccountPrivate(ctx context.Context, accountId primitive.ObjectID) (bool, error) {
	return deleteElemWithId(ctx, c.accountPrivate, accountId)
}
