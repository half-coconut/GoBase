package ioc

import (
	"GoBase/webook/internal/service/memory"
	"GoBase/webook/internal/service/sms"
)

func InitSMSService() sms.Service {
	// 换内存，还是换别的
	return memory.NewService()
	//svc:= memory.NewService()
	//return ratelimit.NewRateLimitSMSService(svc,)
}
