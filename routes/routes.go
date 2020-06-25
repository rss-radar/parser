package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rss-radar/parser/config"
)

func Register(g *gin.Engine) *gin.Engine {
	g.Use(gin.Recovery())
	if config.AppConfig.RunMode != config.RunModeRelease {
		g.Use(gin.Logger())
	}

	g.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	return g
}
