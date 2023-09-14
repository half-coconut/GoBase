package failover

import (
	"GoBase/webook/internal/service/sms"
	"context"
	"errors"
	"log"
	"sync/atomic"
)

type FailoverSMSService struct {
	svcs []sms.Service

	idx uint64
}

func NewFailoverSMSService(svcs []sms.Service) sms.Service {
	return &FailoverSMSService{
		svcs: svcs,
	}
}

func (f FailoverSMSService) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	for _, svc := range f.svcs {
		err := svc.Send(ctx, tplId, args, numbers...)
		// 发送成功
		if err == nil {
			return nil
		}
		//
		log.Println(err)
	}
	return errors.New("全部服务商都失败了")
}

func (f FailoverSMSService) SendV1(ctx context.Context, tplId string, args []string, numbers ...string) error {
	// 我取下一个节点来作为起始节点
	idx := atomic.AddUint64(&f.idx, 1)
	length := uint64(len(f.svcs))
	for i := idx; i < idx+length; i++ {
		svc := f.svcs[int(i%length)]
		err := svc.Send(ctx, tplId, args, numbers...)
		switch err {
		case nil:
			return nil
		case context.DeadlineExceeded, context.Canceled:
			// 超时或者主动取消了
			return err
		default:
			// 输出日志
		}
	}
	return errors.New("全部服务商都失败了")
}