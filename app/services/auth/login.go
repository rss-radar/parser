package oauth

import (
	"errors"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var (
	oauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GH_CLIENT_ID"),
		ClientSecret: os.Getenv("GH_CLIENT_SECRET"),
		Scopes:       []string{"user"},
		Endpoint:     githuboauth.Endpoint,
	}
)

const oauthStateString = "randomstring"

func GithubOauth(state string, code string) (*github.User, error) {
	if state != oauthStateString {
		return nil, errors.New("Invalid oauth state")
	}

	token, err := oauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}

	oauthClient := oauthConfig.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get(oauth2.NoContext, "")
	if err != nil {
		return nil, err
	}

	return user, nil
}

func OAuthCodeURL() string {
	return oauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)
}
