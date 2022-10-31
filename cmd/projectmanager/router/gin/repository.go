package gin

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

// AddRepository godoc
// @Summary add new repository
// @Schemes
// @Description add new repository
// @Accept json
// @Produce json
// @Param data body AddRepositoryRequest true "Repository 등록"
// @Success 200 {object} AddRepositoryResponse
// @Router /api/repository [post]
func (r *Router) AddRepository(c *gin.Context) {
	req := AddRepositoryRequest{}
	resp := AddRepositoryResponse{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	accountId, err := primitive.ObjectIDFromHex(req.AccountId)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	repositoryId, err := r.db.AddRepository(ctx, accountId, req.RepositoryName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	repositoryIdString := repositoryId.Hex()

	resp.RepositoryId = repositoryIdString
	c.JSON(http.StatusOK, resp)
}

// DeleteRepository godoc
// @Summary delete repository
// @Schemes
// @Description delete repository
// @Accept json
// @Produce json
// @Param repositoryId path string true "Repository ID"
// @Success 200 {object} DeleteRepositoryResponse
// @Router /api/repository/delete/{repositoryId} [delete]
func (r *Router) DeleteRepository(c *gin.Context) {
	resp := DeleteRepositoryResponse{}
	repositoryIdString := c.Param("repositoryId")

	repositoryId, err := primitive.ObjectIDFromHex(repositoryIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	notExist, _ := r.db.DeleteRepository(ctx, repositoryId)
	if !notExist {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Message = "ok"
	c.JSON(http.StatusOK, resp)
}

// ListRepository godoc
// @Summary list current repository
// @Schemes
// @Description list current repository
// @Accept json
// @Produce json
// @Success 200 {object} ListRepositoryResponse
// @Router /api/repository [get]
func (r *Router) ListRepository(c *gin.Context) {
	resp := ListRepositoryResponse{}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	repositories, err := r.db.ListRepository(ctx, nil, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Repositories = make([]Repository, len(repositories))
	for index := range resp.Repositories {
		resp.Repositories[index] = Repository{
			Id:        repositories[index].Id.Hex(),
			Name:      repositories[index].Name,
			AccountId: repositories[index].AccountId.Hex(),
		}
	}
	c.JSON(http.StatusOK, resp)
}

// GetRepository godoc
// @Summary get repository info by id
// @Schemes
// @Description get repository info by id
// @Accept json
// @Produce json
// @Param repositoryId path string true "Repository ID"
// @Success 200 {object} GetRepositoryResponse
// @Router /api/repository/detail/{repositoryId} [get]
func (r *Router) GetRepository(c *gin.Context) {
	repositoryIdString := c.Param("repositoryId")
	resp := GetRepositoryResponse{}
	repositoryId, err := primitive.ObjectIDFromHex(repositoryIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	repository, err := r.db.GetRepository(ctx, repositoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Repository = Repository{
		Id:        repository.Id.Hex(),
		Name:      repository.Name,
		AccountId: repository.AccountId.Hex(),
	}
	c.JSON(http.StatusOK, resp)
}
