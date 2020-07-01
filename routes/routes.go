package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/rss-radar/parser/config"
	"github.com/rss-radar/parser/pkg/ginutils/router"
	"github.com/rss-radar/parser/routes/middleware"
)

func Register(g *gin.Engine) *gin.Engine {
	g.Use(gin.Recovery())
	if config.AppConfig.RunMode != config.RunModeRelease {
		g.Use(gin.Logger())
	}

	cookieStore := cookie.NewStore([]byte("secret"))
	g.Use(sessions.Sessions("user_session", cookieStore))

	g.Use(middleware.CurrentUser())

	r := &router.ApplicationRouter{Router: g}
	registerWeb(r)

	return g
}
