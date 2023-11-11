package retryable

import (
	"GoBase/webook/internal/service/sms"
	"context"
	"errors"
)

// Service 小心并发问题
type Service struct {
	svc sms.Service
	// 重试
	retryMax int
}

func (s Service) Send(ctx context.Context, biz string, args []string, numbers ...string) error {
	err := s.svc.Send(ctx, biz, args, numbers...)
	cnt := 1
	for err != nil && cnt < s.retryMax {
		err = s.svc.Send(ctx, biz, args, numbers...)
		if err != nil {
			return err
		}
		cnt++
	}
	return errors.New("重试都失败了")
}
