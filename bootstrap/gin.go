package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/rss-radar/parser/config"
)

func SetupGin(g *gin.Engine) {
	gin.SetMode(config.AppConfig.RunMode)
}
