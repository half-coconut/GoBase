package ratelimit

import (
	"GoBase/webook/internal/service/sms"
	"GoBase/webook/pkg/ratelimit"
	"context"
	"fmt"
)

var errLimiter = fmt.Errorf("触发了限流")

type RateLimitSMSService struct {
	svc     sms.Service
	limiter ratelimit.Limiter
}

func NewRateLimitSMSService(svc sms.Service, limiter ratelimit.Limiter) sms.Service {
	return &RateLimitSMSService{
		svc:     svc,
		limiter: limiter,
	}
}

func (s *RateLimitSMSService) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	limited, err := s.limiter.Limit(ctx, "sms:tencent")
	if err != nil {
		// 系统错误
		// 可以限流，也可以不限流
		return fmt.Errorf("短信服务判断是否限流出现问题，%w", err)
	}
	if limited {
		return errLimiter
	}
	err = s.svc.Send(ctx, tpl, args, numbers...)
	return err
}
