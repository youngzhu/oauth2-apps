// Package basecamp provides constants for using OAuth2 to access the Basecamp API.
package basecamp // import "github.com/youngzhu/oauth2-apps/basecamp"

import "golang.org/x/oauth2"

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://launchpad.37signals.com/authorization/new",
	TokenURL: "https://launchpad.37signals.com/authorization/token",
}
