// Package basecamp provides constants for using OAuth2 to access the Basecamp API.
package basecamp // import "github.com/youngzhu/oauth2-apps/basecamp"

import (
	"context"
	"errors"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"log"
)

const (
	authURL  = "https://launchpad.37signals.com/authorization/new"
	tokenURL = "https://launchpad.37signals.com/authorization/token"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  authURL,
	TokenURL: tokenURL,
}

var Endpoint4Refresh = oauth2.Endpoint{
	AuthURL:  authURL,
	TokenURL: tokenURL + "?type=refresh",
}

// RefreshToken use to request a new access token when it expires (2 week lifetime, currently).
// see https://github.com/basecamp/api/blob/master/sections/authentication.md#get-authorization
// refreshToken - must be refreshToken, NOT the accessToken
func RefreshToken(clientID, clientSecret, refreshToken string) (*oauth2.Token, error) {
	if clientID == "" {
		return nil, errors.New("basecamp: clientID is not set")
	}
	if clientSecret == "" {
		return nil, errors.New("basecamp: clientSecret is not set")
	}
	if refreshToken == "" {
		return nil, errors.New("basecamp: refreshToken is not set")
	}

	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     Endpoint4Refresh,
	}

	return conf.TokenSource(context.Background(), &oauth2.Token{RefreshToken: refreshToken}).Token()
}

func GetAccessToken() (string, bool) {
	refreshToken := viper.GetString("REFRESH_TOKEN")

	if refreshToken != "" {
		clientID := viper.GetString("clientID")
		clientSecret := viper.GetString("clientSecret")

		token, err := RefreshToken(clientID, clientSecret, refreshToken)
		if err == nil {
			return token.AccessToken, true
		} else {
			log.Println("refresh token error:", err)
		}
	}

	return viper.GetString("ACCESS_TOKEN"), false
}
