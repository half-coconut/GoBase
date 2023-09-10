package cache

import (
	"GoBase/webook/internal/repository/cache/redismocks"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestRedisCodeCache_Set(t *testing.T) {
	testCases := []struct {
		name  string
		mock  func(ctrl *gomock.Controller) redis.Cmdable
		ctx   context.Context
		biz   string
		phone string
		code  string

		wantErr error
	}{
		{
			name: "验证码设置成功",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := redismocks.NewMockCmdable(ctrl)
				res := redis.NewCmd(context.Background())
				res.SetVal(int64(0))
				cmd.EXPECT().Eval(gomock.Any(), luaSetCode,
					[]string{"phone_code:login:13800009876"}, []any{"123456"}, // 不定参数，mock 时需要使用切片
				).Return(res)
				return cmd
			},
			ctx:   context.Background(),
			biz:   "login",
			phone: "13800009876",
			code:  "123456",

			wantErr: nil,
		},
		{
			name: "redis 错误",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := redismocks.NewMockCmdable(ctrl)
				res := redis.NewCmd(context.Background())
				res.SetErr(errors.New("mock redis error"))
				cmd.EXPECT().Eval(gomock.Any(), luaSetCode,
					[]string{"phone_code:login:13800009876"}, []any{"123456"}, // 不定参数，mock 时需要使用切片
				).Return(res)
				return cmd
			},
			ctx:   context.Background(),
			biz:   "login",
			phone: "13800009876",
			code:  "123456",

			wantErr: errors.New("mock redis error"),
		},
		{
			name: "发送太频繁",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := redismocks.NewMockCmdable(ctrl)
				res := redis.NewCmd(context.Background())
				res.SetVal(int64(-1))
				cmd.EXPECT().Eval(gomock.Any(), luaSetCode,
					[]string{"phone_code:login:13800009876"}, []any{"123456"}, // 不定参数，mock 时需要使用切片
				).Return(res)
				return cmd
			},
			ctx:   context.Background(),
			biz:   "login",
			phone: "13800009876",
			code:  "123456",

			wantErr: ErrCodeSendTooMany,
		},
		{
			name: "系统错误",
			mock: func(ctrl *gomock.Controller) redis.Cmdable {
				cmd := redismocks.NewMockCmdable(ctrl)
				res := redis.NewCmd(context.Background())
				res.SetVal(int64(-10))
				//res.SetErr(errors.New("系统错误"))
				cmd.EXPECT().Eval(gomock.Any(), luaSetCode,
					[]string{"phone_code:login:13800009876"}, []any{"123456"}, // 不定参数，mock 时需要使用切片
				).Return(res)
				return cmd
			},
			ctx:   context.Background(),
			biz:   "login",
			phone: "13800009876",
			code:  "123456",

			wantErr: ErrUnknownForCode,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			c := NewCodeRedisCache(tc.mock(ctrl))
			err := c.Set(tc.ctx, tc.biz, tc.phone, tc.code)
			assert.Equal(t, tc.wantErr, err)

		})
	}
}
