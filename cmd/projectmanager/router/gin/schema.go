package gin

import "time"

type Message struct {
	Message string `json:"message"`
}

type Account struct {
	Id           string `json:"id"`
	RegistryType string `json:"registry_type"`
	RegistryUrl  string `json:"registry_url"`
	Username     string `json:"username"`
	Nickname     string `json:"nickname"`
}

type AccountPrivate struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type Repository struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	AccountId string `json:"accountId"`
}

type Image struct {
	Id           string    `json:"id"`
	RepositoryId string    `json:"repository_id"`
	Tag          string    `json:"tag"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AddAccountRequest struct {
	RegistryType    string `json:"registry_type" binding:"required"`
	RegistryUrl     string `json:"registry_url"`
	AccountUsername string `json:"account_username" binding:"required"`
	AccountNickname string `json:"account_nickname"`
	AccountPassword string `json:"account_password" binding:"required"`
}

type AddAccountResponse struct {
	AccountId string `json:"account_id"`
}

type UpdateAccountNicknameRequest struct {
	Nickname string `json:"nickname" binding:"required"`
}

type UpdateAccountNicknameResponse Message

type UpdateAccountPasswordRequest struct {
	Password string `json:"password" binding:"required"`
}

type UpdateAccountPasswordResponse Message

type DeleteAccountResponse Message

type GetAccountResponse struct {
	Account Account `json:"account"`
}

type GetAccountByAccountDataRequest struct {
	RegistryType    string `json:"registry_type" binding:"required"`
	RegistryUrl     string `json:"registry_url"`
	AccountUsername string `json:"account_username" binding:"required"`
}

type GetAccountByAccountDataResponse struct {
	Account Account `json:"account"`
}

type GetAccountPrivateResponse struct {
	AccountPrivate AccountPrivate `json:"account_private"`
}

type ListAccountResponse struct {
	Accounts []Account `json:"accounts"`
}

type AddRepositoryRequest struct {
	AccountId      string `json:"account_id" binding:"required"`
	RepositoryName string `json:"repository_name" binding:"required"`
}
type AddRepositoryResponse struct {
	RepositoryId string `json:"repository_id"`
}

type DeleteRepositoryResponse Message

type ListRepositoryResponse struct {
	Repositories []Repository `json:"repositories"`
}

type GetRepositoryRequest struct {
	RepositoryId string `json:"repository_id" binding:"required"`
}

type GetRepositoryResponse struct {
	Repository Repository `json:"repository"`
}

type ListRegistryTypeResponse struct {
	RepositoryTypes []string `json:"repository_types"`
}

type GetImageResponse struct {
	Image Image `json:"image"`
}

type GetLatestImageByRepository struct {
	Image Image `json:"image"`
}

type ListImageByRepository struct {
	Images []Image `json:"images"`
}

type CreateOrUpdateImageRequest struct {
	Tag string `json:"tag" binding:"required"`
}
type CreateOrUpdateImageResponse struct {
	ImageId string `json:"image_id"`
}
