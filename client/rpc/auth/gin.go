package auth

import "github.com/gin-gonic/gin"

// Gin auth Middleware
// 通过 r.Use(<middleware>) 添加中间件
//
//	r := gin.New()
//	r.Use()
func (a *KeyauthAuther) GinAuthHandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
