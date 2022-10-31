package gin

import (
	"github.com/ZeroOneAI/AISecu/cmd/projectmanager/config/viper"
	"github.com/ZeroOneAI/AISecu/cmd/projectmanager/docs"
	"github.com/ZeroOneAI/AISecu/pkg/db/mongodb"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Router struct {
	*gin.Engine
	// TODO db 하나 대신 각 분류별(ex. account, project, 등)로 분리된 인터페이스로 대체 (하나의 Struct 가 여러 Interface 로 저장되어도 상관 X)
	db          *mongodb.Controller
	urlRootPath string
}

const (
	ExternalApiPrefix = "/api"
	InternalApiPrefix = "/internal"
)

func NewRouter(configFilePath string) (*Router, error) {
	var mongoController *mongodb.Controller
	var err error

	config, err := viper.NewFromFile(configFilePath)
	if err != nil {
		return nil, err
	}

	mongoController, err = config.Db.GetController()
	if err != nil {
		return nil, err
	}

	r := &Router{
		Engine:      gin.Default(),
		db:          mongoController,
		urlRootPath: config.ApiRootPath,
	}
	r.db.InitAllCollectionIndex()
	r.setRoute()
	return r, nil
}

func (r *Router) setRoute() {
	// Related to Account
	r.GET(r.urlRootPath+ExternalApiPrefix+"/account/detail/:accountId", r.GetAccount)
	r.GET(r.urlRootPath+ExternalApiPrefix+"/account/list", r.ListAccount)
	r.POST(r.urlRootPath+ExternalApiPrefix+"/account", r.AddAccount)
	r.DELETE(r.urlRootPath+ExternalApiPrefix+"/account/delete/:accountId", r.DeleteAccount)
	r.PUT(r.urlRootPath+ExternalApiPrefix+"/account/nickname/:accountId", r.UpdateAccountNickname)
	r.PUT(r.urlRootPath+ExternalApiPrefix+"/account/password/:accountId", r.UpdateAccountPassword)

	// Related to Registry
	r.GET(r.urlRootPath+ExternalApiPrefix+"/registry/list", r.ListRegistryType)

	// Related to Repository
	r.GET(r.urlRootPath+ExternalApiPrefix+"/repository/detail/:repositoryId", r.GetRepository)
	r.GET(r.urlRootPath+ExternalApiPrefix+"/repository/latest/:repositoryId", r.GetLatestImageByRepository)
	r.PUT(r.urlRootPath+ExternalApiPrefix+"/repository/image/:repositoryId", r.CreateOrUpdateImage)
	r.GET(r.urlRootPath+ExternalApiPrefix+"/repository/images/:repositoryId", r.ListImageByRepository)
	r.GET(r.urlRootPath+ExternalApiPrefix+"/repository", r.ListRepository)
	r.POST(r.urlRootPath+ExternalApiPrefix+"/repository", r.AddRepository)
	r.DELETE(r.urlRootPath+ExternalApiPrefix+"/repository/delete/:repositoryId", r.DeleteRepository)

	// Related to Image
	r.GET(r.urlRootPath+ExternalApiPrefix+"/image/detail/:imageId", r.GetImage)

	// internal API
	r.GET(r.urlRootPath+InternalApiPrefix+"/account/private/:accountId", r.GetAccountPrivate)

	// Init Swagger
	docs.SwaggerInfo.BasePath = r.urlRootPath
	r.GET(r.urlRootPath+"/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
