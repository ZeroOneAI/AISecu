package gin

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

// CreateOrUpdateImage godoc
// @Summary Create or Update Image
// @Schemes
// @Description Create or Update Image
// @Accept json
// @Produce json
// @Param data body CreateOrUpdateImageRequest true "Image 정보"
// @Param repositoryId path string true "Repository ID"
// @Success 200 {object} CreateOrUpdateImageResponse
// @Router /api/repository/image/{repositoryId} [put]
func (r *Router) CreateOrUpdateImage(c *gin.Context) {
	req := CreateOrUpdateImageRequest{}
	resp := CreateOrUpdateImageResponse{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	repositoryIdString := c.Param("repositoryId")
	repositoryId, err := primitive.ObjectIDFromHex(repositoryIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	imageId, err := r.db.CreateOrUpdateImage(ctx, repositoryId, req.Tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.ImageId = imageId.Hex()
	c.JSON(http.StatusOK, resp)
}

// ListImageByRepository godoc
// @Summary List Image By Repository
// @Schemes
// @Description List Image By Repository
// @Accept json
// @Produce json
// @Param repositoryId path string true "Repository ID"
// @Success 200 {object} ListImageByRepository
// @Router /api/repository/images/{repositoryId} [get]
func (r *Router) ListImageByRepository(c *gin.Context) {
	repositoryIdString := c.Param("repositoryId")
	resp := ListImageByRepository{}
	repositoryId, err := primitive.ObjectIDFromHex(repositoryIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	images, err := r.db.ListImageByRepositoryId(ctx, repositoryId, nil, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.Images = make([]Image, len(images))
	for index := range resp.Images {
		resp.Images[index] = Image{
			Id:           images[index].Id.Hex(),
			RepositoryId: images[index].RepositoryId.Hex(),
			Tag:          images[index].Tag,
			UpdatedAt:    images[index].UpdatedAt,
		}
	}
	c.JSON(http.StatusOK, resp)
}

// GetLatestImageByRepository godoc
// @Summary Get Latest Image by Repository
// @Schemes
// @Description Get Latest Image by Repository
// @Accept json
// @Produce json
// @Param repositoryId path string true "Repository ID"
// @Success 200 {object} GetLatestImageByRepository
// @Router /api/repository/latest/{repositoryId} [get]
func (r *Router) GetLatestImageByRepository(c *gin.Context) {
	repositoryIdString := c.Param("repositoryId")
	resp := GetLatestImageByRepository{}
	repositoryId, err := primitive.ObjectIDFromHex(repositoryIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	image, err := r.db.GetLatestImageByRepositoryId(ctx, repositoryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Image = Image{
		Id:           image.Id.Hex(),
		RepositoryId: image.RepositoryId.Hex(),
		Tag:          image.Tag,
		UpdatedAt:    image.UpdatedAt,
	}
	c.JSON(http.StatusOK, resp)
}

func (r *Router) GetImage(c *gin.Context) {
	imageIdString := c.Param("imageId")
	resp := GetImageResponse{}
	imageId, err := primitive.ObjectIDFromHex(imageIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	image, err := r.db.GetImage(ctx, imageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp.Image = Image{
		Id:           image.Id.Hex(),
		RepositoryId: image.RepositoryId.Hex(),
		Tag:          image.Tag,
		UpdatedAt:    image.UpdatedAt,
	}
	c.JSON(http.StatusOK, resp)
}
