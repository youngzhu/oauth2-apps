package basecamp

import (
	"github.com/spf13/viper"
	"testing"
)

func init() {
	viper.SetEnvPrefix("BASECAMP")
	viper.AutomaticEnv() // read in environment variables that match
}

func TestRefreshToken(t *testing.T) {
	clientID := viper.GetString("clientID")
	clientSecret := viper.GetString("clientSecret")
	refreshToken := viper.GetString("REFRESH_TOKEN")

	if clientID == "" {
		t.Error("clientID should not empty")
	}

	if clientSecret == "" {
		t.Error("clientSecret should not empty")
	}

	if refreshToken == "" {
		t.Error("refreshToken should not empty")
	}

	_, err := RefreshToken(clientID, clientSecret, refreshToken)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAccessToken(t *testing.T) {
	_, refresh := GetAccessToken()
	if refresh != true {
		t.Error("should request a new access token by refresh")
	}
}
