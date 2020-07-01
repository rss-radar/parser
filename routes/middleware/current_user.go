package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rss-radar/parser/pkg/auth"
)

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = auth.SaveCurrentUserToContext(c)
		c.Next()
	}

}
