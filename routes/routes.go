package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rss-radar/parser/config"
	"github.com/rss-radar/parser/pkg/ginutils/router"
)

func Register(g *gin.Engine) *gin.Engine {
	g.Use(gin.Recovery())
	if config.AppConfig.RunMode != config.RunModeRelease {
		g.Use(gin.Logger())
	}

	r := &router.ApplicationRouter{Router: g}

	registerWeb(r)

	return g
}
