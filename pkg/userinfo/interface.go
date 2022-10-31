package userinfo

import "github.com/gin-gonic/gin"

type Interface interface {
	GetPassword(accountId string) (string, error)
	SetRepositoryInContextByRepositoryId(con *gin.Context, repositoryId string) error
	SetAccountInContextByAccountId(con *gin.Context, accountId string) error
	RepositoryPushed(repositoryId, tag string) error
}
