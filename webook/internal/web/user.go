package web

import (
	"GoBase/webook/internal/domain"
	"GoBase/webook/internal/service"
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		// 和上面比起来，用 ` 看起来就比较清爽
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		svc:         svc,
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

// RegisterRoutesV1 Group 的第二种处理方式
//func (u *UserHandler) RegisterRoutesV1(ug *gin.RouterGroup) {
//	ug.POST("/signup", u.SignUp)
//	ug.POST("/login", u.Login)
//	ug.POST("/edit", u.Edit)
//	ug.GET("/profile", u.Profile)
//}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
	ug.GET("/profile", u.Profile)
}

func (u *UserHandler) SignUp(c *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}
	var req SignUpReq
	// Bind 方法，会根据 Content-Type 来解析数据到 req
	// 解析错误，会直接写回一个 4xx 的错误
	if err := c.Bind(&req); err != nil {
		return
	}

	//ok, err := regexp.Match(emailRegexPattern, []byte(req.Email))
	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		// 缺点：使用 5xx 或者 4xx，容易区分不出来是是否到达了服务器。
		//c.String(http.StatusInternalServerError, "系统错误 "+http.StatusText(http.StatusInternalServerError))
		// 优点：如果使用 2xx, bizCode 的方式，可以明确知道到达了服务器
		// 缺点：不符合 http 规范，不好做监控
		c.String(http.StatusOK, "系统错误 "+http.StatusText(http.StatusOK))
		return
	}
	if !ok {
		//c.String(http.StatusBadRequest, "邮箱格式不正确 "+http.StatusText(http.StatusBadRequest))
		c.String(http.StatusOK, "邮箱格式不正确 "+http.StatusText(http.StatusOK))
		return
	}
	if req.ConfirmPassword != req.Password {
		c.String(http.StatusOK, "两次输入的密码不一致")
		return
	}
	//ok, err = regexp.Match(passwordRegexPattern, []byte(req.Email))
	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		// 记录日志
		c.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		c.String(http.StatusOK, "密码必须大于 8 位，包含数字、特殊字符")
		return
	}
	err = u.svc.SignUp(c, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.String(http.StatusOK, "系统异常")
		return
	}

	c.String(http.StatusOK, "注册成功！")
	fmt.Printf("%v", req)
	// 这边是数据库的操作

}

func (u *UserHandler) Login(c *gin.Context) {
}

func (u *UserHandler) Edit(c *gin.Context) {

}

func (u *UserHandler) Profile(c *gin.Context) {

}
