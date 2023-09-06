package integration

import (
	"GoBase/webook/internal/web"
	"GoBase/webook/ioc"
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"time"

	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_e2e_SendSMSLoginCode(t *testing.T) {
	server := InitWebServer()
	rdb := ioc.InitRedis()
	testCases := []struct {
		name string
		// 数据准备
		before  func(t *testing.T)
		after   func(t *testing.T)
		reqBody string

		wantCode int
		wantBody web.Result
	}{
		{
			name: "发送成功",
			before: func(t *testing.T) {
				// 不需要
			},
			after: func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
				// 需要清理数据
				val, err := rdb.GetDel(ctx, "phone_code:login:13800001234").Result()
				cancel()
				assert.NoError(t, err)
				assert.True(t, len(val) == 6)
			},
			reqBody:  `{"phone":"13800001234"}`,
			wantCode: http.StatusOK,
			wantBody: web.Result{
				Msg: "发送成功",
			},
		},
		{
			name: "发送太频繁",
			before: func(t *testing.T) {
				// 手机号已经有了验证码了
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
				// 需要清理数据
				_, err := rdb.Set(ctx, "phone_code:login:13800001234", "123456", time.Minute*9+time.Second*30).Result()
				cancel()
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
				// 需要清理数据
				val, err := rdb.GetDel(ctx, "phone_code:login:13800001234").Result()
				cancel()
				assert.NoError(t, err)
				assert.Equal(t, "123456", val)
			},
			reqBody:  `{"phone":"13800001234"}`,
			wantCode: http.StatusOK,
			wantBody: web.Result{
				Code: 4,
				Msg:  "短信发送太频繁，请稍后再试",
			},
		},
		{
			name: "系统错误",
			before: func(t *testing.T) {
				// 手机号已经有了验证码了
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
				// 需要清理数据
				_, err := rdb.Set(ctx, "phone_code:login:13800001234", "123456", 0).Result()
				cancel()
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
				// 需要清理数据
				val, err := rdb.GetDel(ctx, "phone_code:login:13800001234").Result()
				cancel()
				assert.NoError(t, err)
				assert.Equal(t, "123456", val)
			},
			reqBody:  `{"phone":"13800001234"}`,
			wantCode: http.StatusOK,
			wantBody: web.Result{
				Code: 5,
				Msg:  "系统错误",
			},
		},
		{
			name: "手机号码为空",
			before: func(t *testing.T) {
			},
			after: func(t *testing.T) {
			},
			reqBody:  `{"phone":""}`,
			wantCode: http.StatusOK,
			wantBody: web.Result{
				Code: 4,
				Msg:  "请输入手机号码",
			},
		},
		{
			name: "数据格式不对",
			before: func(t *testing.T) {
			},
			after: func(t *testing.T) {
			},
			reqBody:  `{"phone":"}`,
			wantCode: http.StatusBadRequest,
			//wantBody: web.Result{
			//	Code: 4,
			//	Msg:  "请输入手机号码",
			//},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(t)
			req, err := http.NewRequest(http.MethodPost, "/users/login_sms/code/send", bytes.NewBuffer([]byte(tc.reqBody)))
			require.NoError(t, err)
			// 注意 json 数据格式
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			// 这是 http 请求进去，gin 框架的入口
			server.ServeHTTP(resp, req)

			assert.Equal(t, tc.wantCode, resp.Code)
			if resp.Code != 200 {
				return
			}
			var webRes web.Result
			err = json.NewDecoder(resp.Body).Decode(&webRes)
			require.NoError(t, err)
			assert.Equal(t, tc.wantBody, webRes)
			tc.after(t)
		})
	}
}
