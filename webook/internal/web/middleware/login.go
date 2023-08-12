package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
Build 模式，不能对用户调用顺序有任何的假设
*/

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (l *LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginMiddlewareBuilder) Builder() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 不需要登录校验 session
		for _, path := range l.paths {
			if c.Request.URL.Path == path {
				return
			}
		}

		// 登录和注册，不需要登录校验 session
		//if c.Request.URL.Path == "/users/login" || c.Request.URL.Path == "/users/signup" {
		//	return
		//}

		sess := sessions.Default(c)
		id := sess.Get("userId")
		if id == nil {
			// 没有登录
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
