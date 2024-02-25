//go:build wireinject

package startup

import (
	"GoBase/webook/internal/repository"
	"GoBase/webook/internal/repository/cache"
	"GoBase/webook/internal/repository/dao"
	"GoBase/webook/internal/service"
	"GoBase/webook/internal/web"
	"GoBase/webook/ioc"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		ioc.InitDB, ioc.InitRedis, ioc.InitLogger,

		dao.NewUserDAO,
		dao.NewGORMArticleDAO,

		cache.NewUserCache,
		cache.NewCodeRedisCache,

		repository.NewUserRepository,
		repository.NewCodeRepository,
		repository.NewArticleRepository,

		service.NewUserService,
		service.NewCodeService,
		service.NewArticleService,

		ioc.InitSMSService,
		web.NewUserHandler,
		web.NewArticleHandler,

		ioc.InitWebServer,
		ioc.InitMiddlewares,
	)
	return new(gin.Engine)
}

var thirdProvider = wire.NewSet(
	ioc.InitRedis,
	ioc.InitLogger,
	ioc.InitDB,
)

var userSvcProvider = wire.NewSet(
	dao.NewUserDAO,
	cache.NewUserCache,
	cache.NewCodeRedisCache,

	repository.NewUserRepository,
	repository.NewCodeRepository,

	service.NewUserService,
	service.NewCodeService,

	ioc.InitSMSService,
	web.NewUserHandler,
)

func InitArticleHandler() *web.ArticleHandler {
	wire.Build(
		thirdProvider,
		dao.NewGORMArticleDAO,
		repository.NewArticleRepository,
		service.NewArticleService,
		web.NewArticleHandler)
	return &web.ArticleHandler{}
}

func InitUserSvc() service.UserService {
	wire.Build(thirdProvider, userSvcProvider)
	return service.NewUserService(nil)
}

func InitTestDB() service.UserService {
	wire.Build(ioc.InitDB)
	return service.NewUserService(nil)
}
