// Package basecamp provides constants for using OAuth2 to access the Basecamp API.
package basecamp // import "github.com/youngzhu/oauth2-apps/basecamp"

import "golang.org/x/oauth2"

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
