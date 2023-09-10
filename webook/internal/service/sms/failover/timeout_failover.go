package failover

import (
	"GoBase/webook/internal/service/sms"
	"context"
	"errors"
	"sync/atomic"
)

type TimeoutFailoverSMSService struct {
	// 你的服务商
	svcs []sms.Service
	idx  int32
	// 连续超时的个数
	cnt int32

	// 阈值，连续超时，超过了这个数值，就要切换
	threshold int32
}

func (t TimeoutFailoverSMSService) Send(ctx context.Context, tplId string, args []string, numbers ...string) error {
	idx := atomic.LoadInt32(&t.idx)
	cnt := atomic.LoadInt32(&t.cnt)
	if cnt > t.threshold {
		// 这里要切换，新的下标
		newIdx := (idx + 1) % int32(len(t.svcs))
		if atomic.CompareAndSwapInt32(&t.idx, idx, newIdx) {
			// 如果切换成功
			atomic.StoreInt32(&t.cnt, 0)
		} else {
			// 出现并发了
			// 两种写法，都可以
			//idx = newIdx
			idx = atomic.LoadInt32(&t.idx)
		}
		svc := t.svcs[idx]
		err := svc.Send(ctx, tplId, args, numbers...)
		switch err {
		case context.DeadlineExceeded:
			atomic.AddInt32(&t.cnt, 1)
			return err
		case nil:
			// 连续状态被打断
			atomic.StoreInt32(&t.cnt, 0)
			return nil
		default:
			// 不知道什么错误
			// 可以考虑换下一个
			// - 超时可能是偶发的，我尽量再试试
			// - 非超时，我直接下一个
			return err
		}
	}
	return errors.New("全部服务商都失败了")
}

func NewTimeoutFailoverSMSService() sms.Service {

	return nil
}
