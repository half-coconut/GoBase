package web

import (
	"GoBase/webook/internal/domain"
	"GoBase/webook/internal/service"
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	if err == service.ErrUserDuplicateEmail {
		c.String(http.StatusOK, "邮箱冲突")
		return
	}

	if err != nil {
		c.String(http.StatusOK, "系统异常")
		return
	}

	c.String(http.StatusOK, "注册成功！")
	fmt.Printf("%v", req)
	// 这边是数据库的操作

}

func (u *UserHandler) Login(c *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := c.Bind(&req); err != nil {
		return
	}
	user, err := u.svc.Login(c, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		c.String(http.StatusOK, "用户名或密码不正确")
		return
	}
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	// 登录成功后，设置 session
	sess := sessions.Default(c)
	// 要放在 session 里面的东西
	sess.Set("userId", user.Id)
	sess.Save()
	c.String(http.StatusOK, "登录成功！")
	return
}

func (u *UserHandler) Edit(c *gin.Context) {
	type ProfileReq struct {
		Id              int64  `json:"id"`
		NickName        string `json:"nick_name"`
		Birthday        string `json:"birthday"`
		PersonalProfile string `json:"personal_profile"`
	}
	var req ProfileReq
	if err := c.Bind(&req); err != nil {
		return
	}
	if len(req.NickName) >= 50 {
		c.String(http.StatusOK, "昵称名称不超过 50 个字符")
		return
	}
	if len(req.PersonalProfile) >= 200 {
		c.String(http.StatusOK, "个人简介不超过 200 个字符")
		return
	}
	//sess := sessions.Default(c)
	//id := sess.Get("userId")
	//var nid = id.(int64)
	_, err := u.svc.Edit(c, req.Id, req.NickName, req.Birthday, req.PersonalProfile)
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	c.String(http.StatusOK, "编辑 Profile 成功！")
	return
}

func (u *UserHandler) Profile(c *gin.Context) {
	type ProfileReq struct {
		Id int64 `json:"id"`
	}
	var req ProfileReq
	if err := c.Bind(&req); err != nil {
		return
	}
	//sess := sessions.Default(c)
	//id := sess.Get("userId")
	//var nid = id.(int64)
	id := c.Query("id")
	nid, _ := strconv.ParseInt(id, 10, 64)
	user, err := u.svc.Profile(c, nid)
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	c.String(http.StatusOK, "你的昵称是：%v，\n你的生日是：%v，\n你的个人简介：%v", user.NickName, user.Birthday, user.PersonalProfile)
	return
}
