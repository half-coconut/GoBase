package web

import (
	"GoBase/webook/internal/domain"
	"GoBase/webook/internal/service"
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"net/http"
	"strconv"
	"time"
)

const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
	// 和上面比起来，用 ` 看起来就比较清爽
	passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	userIdKey            = "userId"
	bizLogin             = "login"
)

// 确保 UserHandler 上实现了 handler 接口，初始化了一个对象
var _ handler = &UserHandler{}

// 写法二，更优雅，这个没有初始化对象
var _ handler = (*UserHandler)(nil)

type UserHandler struct {
	svc         service.UserService
	codeSvc     service.CodeService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandler(svc service.UserService, codeSvc service.CodeService) *UserHandler {
	return &UserHandler{
		svc:         svc,
		emailExp:    regexp.MustCompile(emailRegexPattern, regexp.None),
		passwordExp: regexp.MustCompile(passwordRegexPattern, regexp.None),
		codeSvc:     codeSvc,
	}
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.POST("/signup", u.SignUp)
	//ug.POST("/login", u.Login)
	ug.POST("/login", u.LoginJWT)
	ug.POST("/edit", u.Edit)
	//ug.GET("/profile", u.Profile)
	ug.GET("/profile", u.ProfileJWT)
	ug.POST("/login_sms/code/send", u.SendSMSLoginCode)
	ug.POST("/login_sms", u.LoginSMS)
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
		c.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		//c.String(http.StatusBadRequest, "邮箱格式不正确 "+http.StatusText(http.StatusBadRequest))
		c.String(http.StatusOK, "邮箱格式不正确")
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
	if err == service.ErrUserDuplicate {
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

func (u *UserHandler) LoginJWT(c *gin.Context) {
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
	// 登录成功后，用 JWT 设置登录态
	// 生成一个 JWT token

	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
		},
		Uid:       user.Id,
		UserAgent: c.Request.UserAgent(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err := token.SignedString([]byte("iyI1vQON0NmwDnaOMZAgdcJQZ7N6TYbD"))
	if err != nil {
		c.String(http.StatusInternalServerError, "系统错误")
		return
	}
	fmt.Println(tokenStr)
	fmt.Println(user)
	c.Header("x-jwt-token", tokenStr)
	c.String(http.StatusOK, "登录成功！")
	return
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
	sess.Options(sessions.Options{
		//Secure: true,
		//HttpOnly: true,
		MaxAge: 60 * 30, // 过期时间: 30min
	})
	sess.Save()
	c.String(http.StatusOK, "登录成功！")
	return
}

func (u *UserHandler) Logout(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Options(sessions.Options{
		//Secure: true,
		//HttpOnly: true,
		MaxAge: -1,
	})
	sess.Save()
	c.String(http.StatusOK, "退出登录成功！")
	return
}

func (u *UserHandler) Edit(ctx *gin.Context) {
	type Req struct {
		// 注意，其它字段，尤其是密码、邮箱和手机，
		// 修改都要通过别的手段
		// 邮箱和手机都要验证
		// 密码更加不用多说了
		Id              int64  `json:"id"`
		NickName        string `json:"nick_name"`
		Birthday        string `json:"birthday"`
		PersonalProfile string `json:"personal_profile"`
	}

	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	// 你可以尝试在这里校验。
	// 比如说你可以要求 Nickname 必须不为空
	// 校验规则取决于产品经理
	if req.NickName == "" {
		ctx.JSON(http.StatusOK, Result{Code: 4, Msg: "昵称不能为空"})
		return
	}

	if len(req.PersonalProfile) > 1024 {
		ctx.JSON(http.StatusOK, Result{Code: 4, Msg: "关于我过长"})
		return
	}
	birthday, err := time.Parse(time.DateOnly, req.Birthday)
	if err != nil {
		// 也就是说，我们其实并没有直接校验具体的格式
		// 而是如果你能转化过来，那就说明没问题
		ctx.JSON(http.StatusOK, Result{Code: 4, Msg: "日期格式不对"})
		return
	}

	uc := ctx.MustGet("user").(UserClaims)
	err = u.svc.UpdateNonSensitiveInfo(ctx, domain.User{
		Id:              uc.Uid,
		NickName:        req.NickName,
		PersonalProfile: req.PersonalProfile,
		Birthday:        birthday,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, Result{Code: 5, Msg: "系统错误"})
		return
	}
	ctx.JSON(http.StatusOK, Result{Msg: "OK"})
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
func (u *UserHandler) ProfileJWT(c *gin.Context) {
	type ProfileReq struct {
		Id int64 `json:"id"`
	}
	var req ProfileReq
	if err := c.Bind(&req); err != nil {
		return
	}
	cl, ok := c.Get("claims")
	if !ok {
		// 监控这里，建议保留
		c.String(http.StatusOK, "系统错误")
		return
	}
	// cl.(*UserClaims) 类型断言,ok 代表是不是 *UserClaims
	claims, ok := cl.(*UserClaims)
	if !ok {
		// 监控这里
		c.String(http.StatusOK, "系统错误")
		return
	}

	user, err := u.svc.Profile(c, claims.Uid)
	if err != nil {
		c.String(http.StatusOK, "系统错误")
		return
	}
	c.String(http.StatusOK, "你的昵称是：%v，\n你的生日是：%v，\n你的个人简介：%v", user.NickName, user.Birthday, user.PersonalProfile)
	return
}

// SendSMSLoginCode 发送短信验证码
func (u *UserHandler) SendSMSLoginCode(ctx *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	// 你也可以用正则表达式校验是不是合法的手机号
	if req.Phone == "" {
		ctx.JSON(http.StatusOK, Result{Code: 4, Msg: "请输入手机号码"})
		return
	}
	err := u.codeSvc.Send(ctx, bizLogin, req.Phone)
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, Result{Msg: "发送成功"})
	case service.ErrCodeSendTooMany:
		ctx.JSON(http.StatusOK, Result{Code: 4, Msg: "短信发送太频繁，请稍后再试"})
	default:
		ctx.JSON(http.StatusOK, Result{Code: 5, Msg: "系统错误"})
		// 要打印日志
		return
	}
}
func (u *UserHandler) LoginSMS(ctx *gin.Context) {
	type Req struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	ok, err := u.codeSvc.Verify(ctx, bizLogin, req.Phone, req.Code)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 5, Msg: "系统异常",
		})
		return
	}
	if !ok {
		ctx.JSON(http.StatusOK, Result{
			Code: 4, Msg: "验证码错误",
		})
		return
	}

	// 验证码是对的
	// 登录或者注册用户
	ue, err := u.svc.FindOrCreate(ctx, req.Phone)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Code: 4, Msg: "系统错误",
		})
		return
	}
	err = u.setJWTToken(ctx, ue.Id)
	if err != nil {
		ctx.JSON(http.StatusOK, Result{
			Msg: "系统错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, Result{Msg: "登录成功"})
}
func (u *UserHandler) setJWTToken(ctx *gin.Context, uid int64) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		Uid:       uid,
		UserAgent: ctx.GetHeader("User-Agent"),
		RegisteredClaims: jwt.RegisteredClaims{
			// 演示目的设置为一分钟过期
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
			// 在压测的时候，要将过期时间设置更长一些
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
		},
	})
	tokenStr, err := token.SignedString(JWTKey)
	if err != nil {
		return err
	}
	ctx.Header("x-jwt-token", tokenStr)
	return nil
}

type UserClaims struct {
	jwt.RegisteredClaims
	// 申明要放进 token 里面的数据
	Uid       int64
	UserAgent string
}

// JWTKey 因为 JWT Key 不太可能变，所以可以直接写成常量
// 也可以考虑做成依赖注入
var JWTKey = []byte("moyn8y9abnd7q4zkq2m73yw8tu9j5ixm")
