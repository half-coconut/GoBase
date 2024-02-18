//go:build wireinject

package startuo

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

		cache.NewUserCache,
		cache.NewCodeRedisCache,

		repository.NewUserRepository,
		repository.NewCodeRepository,

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
