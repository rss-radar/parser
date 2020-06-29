package bootstrap

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"github.com/rss-radar/parser/config"
)

func SetupGin(g *gin.Engine) {
	gin.SetMode(config.AppConfig.RunMode)

	g.HTMLRender = ginview.Default()
}
