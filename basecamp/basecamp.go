// Package basecamp provides constants for using OAuth2 to access the Basecamp API.
package basecamp // import "github.com/youngzhu/oauth2-apps/basecamp"

import (
	"context"
	"golang.org/x/oauth2"
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
func RefreshToken(clientID, clientSecret, refreshToken string) error {
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     Endpoint4Refresh,
		// 根据各网站的API的权限范围设置scopes
		// 不确定时，可以不设置
		Scopes: []string{""},
	}

	_, err := conf.TokenSource(context.Background(), &oauth2.Token{RefreshToken: refreshToken}).Token()

	return err
}
