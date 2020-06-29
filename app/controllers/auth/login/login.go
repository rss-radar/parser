package login

import (
	"github.com/gin-gonic/gin"
	"github.com/rss-radar/parser/app/controllers"
)

func Show(c *gin.Context) {
	controllers.Render(c, "auth/login/show", gin.H{})
}
