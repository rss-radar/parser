package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type renderData = map[string]interface{}

func Render(c *gin.Context, templatePath string, data renderData) {
	c.HTML(http.StatusOK, templatePath, data)
}
