package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/rss-radar/parser/routes"
)

func SetupRouter(g *gin.Engine) {
	routes.Register(g)
}
