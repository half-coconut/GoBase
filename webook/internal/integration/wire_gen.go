// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package integration

import (
	"GoBase/webook/internal/repository"
	"GoBase/webook/internal/repository/cache"
	"GoBase/webook/internal/repository/dao"
	"GoBase/webook/internal/service"
	"GoBase/webook/internal/web"
	"GoBase/webook/ioc"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func InitWebServer() *gin.Engine {
	cmdable := ioc.InitRedis()
	v := ioc.InitMiddlewares(cmdable)
	db := ioc.InitDB()
	userDAO := dao.NewUserDAO(db)
	userCache := cache.NewUserCache(cmdable)
	userRepository := repository.NewUserRepository(userDAO, userCache)
	userService := service.NewUserService(userRepository)
	codeCache := cache.NewCodeRedisCache(cmdable)
	codeRepository := repository.NewCodeRepository(codeCache)
	smsService := ioc.InitSMSService()
	codeService := service.NewCodeService(codeRepository, smsService)
	userHandler := web.NewUserHandler(userService, codeService)
	engine := ioc.InitWebServer(v, userHandler)
	return engine
}
