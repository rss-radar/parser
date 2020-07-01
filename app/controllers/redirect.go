package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Redirect(c *gin.Context, path string) {
	c.Redirect(http.StatusTemporaryRedirect, path)
}
