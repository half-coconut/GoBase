package ratelimit

import "context"

type Limiter interface {
	// Limit 有没有触发限流，key 是限流对象，bool 表示是否限流
	Limit(ctx context.Context, key string) (bool, error)
}
