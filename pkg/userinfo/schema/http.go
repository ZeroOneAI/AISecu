package schema

type AddAccountRequest struct {
	RegistryUrl string `json:"registry_url" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type AddAccountResponse Msg

type AddRepositoryRequest struct {
	RegistryUrl    string `json:"registry_url" binding:"required"`
	Username       string `json:"username" binding:"required"`
	RepositoryName string `json:"repository_name" binding:"required"`
}

type AddRepositoryResponse Msg

type Msg struct {
	Message string `json:"message" binding:"required"`
}
