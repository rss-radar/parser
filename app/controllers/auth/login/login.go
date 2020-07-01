package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rss-radar/parser/app/controllers"
	userModel "github.com/rss-radar/parser/app/models/user"
	oauth "github.com/rss-radar/parser/app/services/auth"
	"github.com/rss-radar/parser/pkg/auth"
)

func LoginHandler(c *gin.Context) {
	controllers.Redirect(c, oauth.OAuthCodeURL())
}

func CallbackHandler(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")
	user, err := oauth.GithubOauth(state, code)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var currentUser *userModel.User

	existUser, err := userModel.FindByUsername(*user.Login)
	if existUser.Username == "" {
		newUser := &userModel.User{Username: *user.Login}
		if err := newUser.Create(); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		currentUser = newUser

	} else {
		currentUser = existUser
	}

	auth.Login(c, currentUser)

	controllers.Redirect(c, "/")
}
