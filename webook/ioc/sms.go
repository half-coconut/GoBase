package ioc

import (
	"GoBase/webook/internal/service/memory"
	"GoBase/webook/internal/service/sms"
)

func InitSMSService() sms.Service {
	return memory.NewService()
}
