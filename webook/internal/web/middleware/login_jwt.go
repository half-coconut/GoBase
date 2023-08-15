package middleware

import (
	"GoBase/webook/internal/web"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
)

/**
Build 模式，不能对用户调用顺序有任何的假设
*/

// LoginJWTMiddlewareBuilder JWT 登录校验
type LoginJWTMiddlewareBuilder struct {
	paths []string
}

func NewLoginJWTMiddlewareBuilder() *LoginJWTMiddlewareBuilder {
	return &LoginJWTMiddlewareBuilder{}
}

func (l *LoginJWTMiddlewareBuilder) IgnorePaths(path string) *LoginJWTMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginJWTMiddlewareBuilder) Builder() gin.HandlerFunc {
	// 用 go 的方式，编码解码为二进制
	gob.Register(time.Now())
	return func(c *gin.Context) {
		// 不需要登录校验 session
		for _, path := range l.paths {
			if c.Request.URL.Path == path {
				return
			}
		}
		// 使用 JWT 校验
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			// 没登录
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		segs := strings.Split(tokenHeader, " ")
		if len(segs) != 2 {
			// 没登录
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := segs[1]
		claims := &web.UserClaims{}
		// ParseWithClaims 里面一定要传入指针
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("iyI1vQON0NmwDnaOMZAgdcJQZ7N6TYbD"), nil
		})
		if err != nil {
			// 没登录
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//err 为 nil，token 不为 nil
		if token == nil || !token.Valid || claims.Uid == 0 {
			// 没登录
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if claims.UserAgent != c.Request.UserAgent() {
			// 严重的安全问题，需要加监控
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		now := time.Now()
		// 每十秒刷新一次
		if claims.ExpiresAt.Sub(now) < time.Second*50 {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))
			tokenStr, err := token.SignedString([]byte("iyI1vQON0NmwDnaOMZAgdcJQZ7N6TYbD"))
			if err != nil {
				log.Println("jwt 续约失败", err)
			}
			c.Header("x-jwt-token", tokenStr)
		}

		c.Set("claims", claims)
	}
}
