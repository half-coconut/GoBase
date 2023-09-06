package web

import (
	"GoBase/webook/internal/domain"
	"GoBase/webook/internal/service"
	svcmocks "GoBase/webook/internal/service/mocks"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEncrypto(t *testing.T) {
	password := "123456"
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatal(err)
	}
	err = bcrypt.CompareHashAndPassword(encrypted, []byte(password))
	assert.NoError(t, err)
}

func TestNil(t *testing.T) {
	testTypeAssert(nil)

}

func testTypeAssert(c any) {
	claims := c.(*UserClaims)
	println(claims.Uid)
}

func TestUserHandler_SignUp(t *testing.T) {
	testCases := []struct {
		name     string
		mock     func(ctrl *gomock.Controller) service.UserService
		reqBody  string
		wantCode int
		wantBody string
	}{
		{
			name: "注册成功",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				usersvc.EXPECT().SignUp(gomock.Any(), domain.User{
					Email:    "123@qq.com",
					Password: "12345#hello",
				}).Return(nil)
				return usersvc
			},
			reqBody: `{
"email":"123@qq.com",
"confirmPassword":"12345#hello",
"password":"12345#hello"
}`,
			wantCode: http.StatusOK,
			wantBody: "注册成功！",
		},
		{
			name: "参数不对，band 失败",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				return usersvc
			},
			reqBody: `{
"email":"123@qq.com",
"password":'12345#hello'
}`,
			wantCode: http.StatusBadRequest,
		},
		{
			name: "邮箱格式不正确",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				return usersvc
			},
			reqBody: `{
"email":"123@q",
"confirmPassword":"12345#hello",
"password":"12345#hello"
}`,
			wantCode: http.StatusOK,
			wantBody: "邮箱格式不正确",
		},
		{
			name: "两次输入的密码不一致",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				return usersvc
			},
			reqBody: `{
"email":"123@qq.com",
"password":"12345#hello",
"confirmPassword":""
}`,
			wantCode: http.StatusOK,
			wantBody: "两次输入的密码不一致",
		},
		{
			name: "密码必须大于 8 位，包含数字、特殊字符",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				return usersvc
			},
			reqBody: `{
"email":"123@qq.com",
"confirmPassword":"12345",
"password":"12345"
}`,
			wantCode: http.StatusOK,
			wantBody: "密码必须大于 8 位，包含数字、特殊字符",
		},
		{
			name: "邮箱冲突",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				usersvc.EXPECT().SignUp(gomock.Any(), domain.User{
					Email:    "123@qq.com",
					Password: "12345#hello",
				}).Return(service.ErrUserDuplicate)
				return usersvc
			},
			reqBody: `{
"email":"123@qq.com",
"confirmPassword":"12345#hello",
"password":"12345#hello"
}`,
			wantCode: http.StatusOK,
			wantBody: "邮箱冲突",
		},
		{
			name: "系统异常",
			mock: func(ctrl *gomock.Controller) service.UserService {
				usersvc := svcmocks.NewMockUserService(ctrl)
				usersvc.EXPECT().SignUp(gomock.Any(), domain.User{
					Email:    "123@qq.com",
					Password: "12345#hello",
				}).Return(errors.New("system error"))
				return usersvc
			},
			reqBody: `{
"email":"123@qq.com",
"confirmPassword":"12345#hello",
"password":"12345#hello"
}`,
			wantCode: http.StatusOK,
			wantBody: "系统异常",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			server := gin.Default()

			h := NewUserHandler(tc.mock(ctrl), nil)
			h.RegisterRoutes(server)

			req, err := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewBuffer([]byte(tc.reqBody)))
			require.NoError(t, err)
			// 注意 json 数据格式
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			println(resp, req)
			//resp.Code

			// 这是 http 请求进去，gin 框架的入口
			server.ServeHTTP(resp, req)

			assert.Equal(t, tc.wantCode, resp.Code)
			assert.Equal(t, tc.wantBody, resp.Body.String())
		})
	}
}

func TestMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	usersvc := svcmocks.NewMockUserService(ctrl)
	// Times: 可以调用的次数
	usersvc.EXPECT().SignUp(gomock.Any(), gomock.Any()).Times(1).
		Return(errors.New("mock error"))

	// context.Background() 创建一个空白的上下文
	err := usersvc.SignUp(context.Background(), domain.User{
		Email: "789@qq.com",
	})
	t.Log(err)
}

func TestUserHandler_LoginSMS(t *testing.T) {
	testCases := []struct {
		name string
		mock func(ctrl *gomock.Controller) (service.UserService, service.CodeService)

		reqBody      string
		wantHttpCode int
		wantMsg      string
		wantCode     int
	}{
		{
			name: "登录成功，用户已存在",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				codesvc := svcmocks.NewMockCodeService(ctrl)
				usersvc := svcmocks.NewMockUserService(ctrl)
				codesvc.EXPECT().Verify(
					gomock.Any(), "login", "138", "123456").
					Return(true, nil)
				usersvc.EXPECT().FindOrCreate(gomock.Any(), "138").
					Return(domain.User{
						Phone: "138",
					}, nil)
				return usersvc, codesvc
			},
			reqBody: `{
"phone":"138",
"code":"123456"
}`,
			wantMsg:      "登录成功",
			wantCode:     0,
			wantHttpCode: http.StatusOK,
		},
		{
			name: "入参格式不对",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				codesvc := svcmocks.NewMockCodeService(ctrl)
				usersvc := svcmocks.NewMockUserService(ctrl)
				return usersvc, codesvc
			},
			reqBody: `{
"phone":"138",
"code":"123
}`,
			wantMsg:      "",
			wantCode:     0,
			wantHttpCode: http.StatusBadRequest,
		},
		{
			name: "code 校验时，系统异常",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				codesvc := svcmocks.NewMockCodeService(ctrl)
				usersvc := svcmocks.NewMockUserService(ctrl)
				codesvc.EXPECT().Verify(
					gomock.Any(), "login", "138", "123456").
					Return(true, errors.New("code verify error"))
				return usersvc, codesvc
			},
			reqBody: `{
"phone":"138",
"code":"123456"
}`,
			wantMsg:      "系统异常",
			wantCode:     5,
			wantHttpCode: http.StatusOK,
		},
		{
			name: "code 校验时，验证码错误",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				codesvc := svcmocks.NewMockCodeService(ctrl)
				usersvc := svcmocks.NewMockUserService(ctrl)
				codesvc.EXPECT().Verify(
					gomock.Any(), "login", "138", "123456").
					Return(false, nil)
				return usersvc, codesvc
			},
			reqBody: `{
"phone":"138",
"code":"123456"
}`,
			wantMsg:      "验证码错误",
			wantCode:     4,
			wantHttpCode: http.StatusOK,
		},
		{
			name: "数据库异常",
			mock: func(ctrl *gomock.Controller) (service.UserService, service.CodeService) {
				codesvc := svcmocks.NewMockCodeService(ctrl)
				usersvc := svcmocks.NewMockUserService(ctrl)
				codesvc.EXPECT().Verify(
					gomock.Any(), "login", "138", "123456").
					Return(true, nil)
				usersvc.EXPECT().FindOrCreate(gomock.Any(), "138").
					Return(domain.User{}, errors.New("mock db error"))
				return usersvc, codesvc
			},
			reqBody: `{
"phone":"138",
"code":"123456"
}`,
			wantMsg:      "系统错误",
			wantCode:     4,
			wantHttpCode: http.StatusOK,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			server := gin.Default()
			h := NewUserHandler(tc.mock(ctrl))
			h.RegisterRoutes(server)

			req, err := http.NewRequest(http.MethodPost, "/users/login_sms", bytes.NewBuffer([]byte(tc.reqBody)))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			//server.Use(func(c *gin.Context) {
			//	c.Set("user", UserClaims{
			//		Uid: 1,
			//	})
			//})
			println(req, resp)
			server.ServeHTTP(resp, req)
			data, _ := stringToJson(resp.Body.String())
			println(resp.Body.String())
			assert.Equal(t, tc.wantHttpCode, resp.Code)
			assert.Equal(t, tc.wantMsg, data.Msg)
			assert.Equal(t, tc.wantCode, data.Code)
		})
	}
}

func stringToJson(res string) (data Result, err error) {
	err = json.Unmarshal([]byte(res), &data)
	if err != nil {
		return Result{}, err
	}
	return data, nil
}
