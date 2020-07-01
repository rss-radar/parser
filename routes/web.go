package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rss-radar/parser/app/controllers/auth/login"
	userModel "github.com/rss-radar/parser/app/models/user"
	"github.com/rss-radar/parser/pkg/ginutils/router"
)

func registerWeb(r *router.ApplicationRouter, middlewares ...gin.HandlerFunc) {
	r.Middleware(middlewares...)

	r.Register("GET", "/", func(c *gin.Context) {
		currentUser := c.Keys["current_user"].(*userModel.User)
		c.JSON(200, gin.H{
			"current_user": currentUser.Username,
		})
	})
	r.Register("GET", "/login", login.LoginHandler)
	r.Register("GET", "/auth/callback", login.CallbackHandler)
}
