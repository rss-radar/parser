package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rss-radar/parser/app/controllers/auth/login"
	"github.com/rss-radar/parser/pkg/ginutils/router"
)

func registerWeb(r *router.ApplicationRouter, middlewares ...gin.HandlerFunc) {
	r.Middleware(middlewares...)

	r.Register("GET", "/login", login.Show)
}
