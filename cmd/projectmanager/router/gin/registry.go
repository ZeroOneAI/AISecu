package gin

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// ListRegistryType godoc
// @Summary list registry type
// @Schemes
// @Description list registry type
// @Accept json
// @Produce json
// @Success 200 {object} ListRegistryTypeResponse
// @Router /api/registry/list [get]
func (r *Router) ListRegistryType(c *gin.Context) {
	resp := ListRegistryTypeResponse{}
	/* request body 가 있을 시 아래 코드 주석 해제
	req := ListRegistryTypeRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	*/

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	regitryTypes, err := r.db.ListRegistryType(ctx, nil, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.RepositoryTypes = regitryTypes
	c.JSON(http.StatusOK, resp)
}
