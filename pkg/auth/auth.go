package auth

import (
	"errors"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	userModel "github.com/rss-radar/parser/app/models/user"
)

func Login(c *gin.Context, user *userModel.User) {
	session := sessions.Default(c)
	session.Set("user", user.Username)
	fmt.Printf("%s", user.Username)
	session.Save()
}

func GetCurrentUserFromContext(c *gin.Context) (*userModel.User, error) {
	err := errors.New("Invalid context")

	userDataFromContext := c.Keys["current_user"]
	if userDataFromContext == nil {
		return nil, err
	}

	user, ok := userDataFromContext.(*userModel.User)
	if !ok {
		return nil, err
	}

	return user, nil
}

func SaveCurrentUserToContext(c *gin.Context) *userModel.User {
	user, err := getCurrentUserFromSession(c)
	if err != nil {
		return nil
	}
	c.Keys["current_user"] = user

	return user
}

func getCurrentUserFromSession(c *gin.Context) (*userModel.User, error) {
	session := sessions.Default(c)
	username := session.Get("user")
	if username == nil {
		return nil, errors.New("Invalid session")
	}

	user, err := userModel.FindByUsername(username.(string))
	if err != nil {
		return nil, err
	}

	return user, nil
}
