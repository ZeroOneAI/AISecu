package mongodb

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"testing"
)

func TestController_Account(t *testing.T) {
	art := assert.New(t)

	con, err := NewController("mongodb", "localhost", "27017", "manager_test", nil, "")
	if err != nil {
		art.NoError(err)
		return
	}
	err = con.InitAllCollectionIndex()
	art.NoError(err)
	testCnt := 10
	i := 0
	// TODO deferf 를 for 문 밖으로 빼내기
	for i = 0; i < testCnt/2; i++ {
		accountId, err := con.AddAccount(context.Background(), "dockerhub", "docker.io", strconv.Itoa(i), "", "asdf")
		if err != nil {
			art.NoError(err)
			return
		}
		// TODO remove defer in loop
		defer delTest(con, accountId, art)
	}
	for ; i < testCnt; i++ {
		accountId, err := con.AddAccount(context.Background(), "dockerhub", "docker.io", strconv.Itoa(i), "nnickname", "asdf")
		if err != nil {
			art.NoError(err)
			return
		}
		// TODO remove defer in loop
		defer delTest(con, accountId, art)
	}
}

func delTest(con *Controller, accountId primitive.ObjectID, art *assert.Assertions) {
	noExist, err := con.DeleteAccount(context.Background(), accountId)
	if err != nil {
		art.NoError(err)
		return
	}
	art.Equal(true, noExist)
}
