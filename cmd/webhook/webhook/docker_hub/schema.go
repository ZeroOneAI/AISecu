package docker_hub

type PayLoad struct {
	CallbackUrl string         `json:"callback_url"`
	PushData    PushDataInfo   `json:"push_data"`
	Repository  RepositoryInfo `json:"repository"`
}

type PushDataInfo struct {
	PushedAt uint64 `json:"pushed_at"`
	Pusher   string `json:"pusher"`
	Tag      string `json:"tag"`
}

type RepositoryInfo struct {
	CommentCount    uint64 `json:"comment_count"`
	DateCreated     uint64 `json:"date_created"`
	Description     string `json:"description"`
	Dockerfile      string `json:"dockerfile"`
	FullDescription string `json:"full_description"`
	IsOfficial      bool   `json:"is_official"`
	IsPrivate       bool   `json:"is_private"`
	IsTrusted       bool   `json:"is_trusted"`
	Name            string `json:"name"`
	Namespace       string `json:"namespace"`
	Owner           string `json:"owner"`
	RepoName        string `json:"repo_name"`
	RepoUrl         string `json:"repo_url"`
	StarCount       uint64 `json:"star_count"`
	Status          string `json:"status"`
}
