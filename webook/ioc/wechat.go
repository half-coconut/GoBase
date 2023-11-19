package ioc

import (
	"GoBase/webook/internal/service/oauth2/wechat"
	"os"
)

func InitWechatService() wechat.Service {
	os.Setenv("WECHAT_APP_ID", "1111")
	os.Setenv("WECHAT_APP_SECRET", "2222")
	appId, ok := os.LookupEnv("WECHAT_APP_ID")
	if !ok {
		panic("没有找到环境变量 WECHAT_APP_ID")
	}
	appKey, ok := os.LookupEnv("WECHAT_APP_SECRET")
	if !ok {
		panic("没有找到环境变量 WECHAT_APP_SECRET")
	}
	return wechat.NewService(appId, appKey)
}
