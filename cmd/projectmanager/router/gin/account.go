package gin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

// AddAccount godoc
// @Summary add new account
// @Schemes
// @Description add new account
// @Accept json
// @Produce json
// @Param data body AddAccountRequest true "새로운 Account 등록"
// @Success 200 {object} AddAccountResponse
// @Router /api/account [post]
func (r *Router) AddAccount(c *gin.Context) {
	req := AddAccountRequest{}
	resp := AddAccountResponse{}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	accountId, err := r.db.AddAccount(ctx, req.RegistryType, req.RegistryUrl, req.AccountUsername, req.AccountNickname, req.AccountPassword)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	accountIdString := accountId.Hex()

	resp.AccountId = accountIdString
	c.JSON(http.StatusOK, resp)
}

// UpdateAccountNickname godoc
// @Summary update nickname of existing account
// @Schemes
// @Description update nickname of existing account
// @Accept json
// @Produce json
// @Param data body UpdateAccountNicknameRequest true "기존 Account Nickname 수정"
// @Param accountId path string true "Account ID"
// @Success 200 {object} UpdateAccountNicknameResponse
// @Router /api/account/nickname/{accountId} [put]
func (r *Router) UpdateAccountNickname(c *gin.Context) {
	req := UpdateAccountNicknameRequest{}
	resp := UpdateAccountNicknameResponse{}
	accountIdString := c.Param("accountId")
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	accountId, err := primitive.ObjectIDFromHex(accountIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	err = r.db.UpdateAccountNickname(ctx, accountId, req.Nickname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Message = "ok"
	c.JSON(http.StatusOK, resp)
}

// UpdateAccountPassword godoc
// @Summary update password of existing account
// @Schemes
// @Description update password of existing account
// @Accept json
// @Produce json
// @Param data body UpdateAccountPasswordRequest true "기존 Account Password 수정"
// @Param accountId path string true "Account ID"
// @Success 200 {object} UpdateAccountPasswordResponse
// @Router /api/account/password/{accountId} [put]
func (r *Router) UpdateAccountPassword(c *gin.Context) {
	req := UpdateAccountPasswordRequest{}
	resp := UpdateAccountPasswordResponse{}
	accountIdString := c.Param("accountId")
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	accountId, err := primitive.ObjectIDFromHex(accountIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	err = r.db.UpdateAccountPassword(ctx, accountId, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Message = "ok"
	c.JSON(http.StatusOK, resp)
}

// DeleteAccount godoc
// @Summary delete account
// @Schemes
// @Description delete account
// @Accept json
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {object} DeleteAccountResponse
// @Router /api/account/delete/{accountId} [delete]
func (r *Router) DeleteAccount(c *gin.Context) {
	resp := DeleteAccountResponse{}
	accountIdString := c.Param("accountId")

	accountId, err := primitive.ObjectIDFromHex(accountIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	notExist, _ := r.db.DeleteAccount(ctx, accountId)
	if !notExist {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Message = "ok"
	c.JSON(http.StatusOK, resp)
}

// GetAccount godoc
// @Summary get account info by id
// @Schemes
// @Description get account info by id
// @Accept json
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {object} GetAccountResponse
// @Router /api/account/detail/{accountId} [get]
func (r *Router) GetAccount(c *gin.Context) {
	accountIdString := c.Param("accountId")
	resp := GetAccountResponse{}
	accountId, err := primitive.ObjectIDFromHex(accountIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	account, err := r.db.GetAccount(ctx, accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Account = Account{
		Id:           account.Id.Hex(),
		RegistryType: account.RegistryType,
		RegistryUrl:  account.RegistryUrl,
		Username:     account.Username,
		Nickname:     account.Nickname,
	}
	c.JSON(http.StatusOK, resp)
}

// GetAccountPrivate godoc
// @Summary get account info by id
// @Schemes
// @Description get account info by id
// @Accept json
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {object} GetAccountPrivateResponse
// @Router /internal/account/{accountId}/private [get]
func (r *Router) GetAccountPrivate(c *gin.Context) {
	resp := GetAccountPrivateResponse{}

	accountIdString := c.Param("accountId")
	accountId, err := primitive.ObjectIDFromHex(accountIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	accountPrivate, err := r.db.GetAccountPrivate(ctx, accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.AccountPrivate = AccountPrivate{
		Id:       accountPrivate.AccountId.Hex(),
		Password: accountPrivate.Password,
	}
	c.JSON(http.StatusOK, resp)
}

// ListAccount godoc
// @Summary list current accounts
// @Schemes
// @Description list current accounts
// @Accept json
// @Produce json
// @Success 200 {object} ListAccountResponse
// @Router /api/account/list [get]
func (r *Router) ListAccount(c *gin.Context) {
	resp := ListAccountResponse{}

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	accounts, err := r.db.ListAccount(ctx, nil, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	resp.Accounts = make([]Account, len(accounts))
	for index := range resp.Accounts {
		resp.Accounts[index] = Account{
			Id:           accounts[index].Id.Hex(),
			RegistryType: accounts[index].RegistryType,
			RegistryUrl:  accounts[index].RegistryUrl,
			Username:     accounts[index].Username,
			Nickname:     accounts[index].Nickname,
		}
	}
	c.JSON(http.StatusOK, resp)
}
