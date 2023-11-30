package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/youngzhu/oauth2-apps/basecamp"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

var clientID, clientSecret string
var conf *oauth2.Config

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	clientID = viper.GetString("clientID")
	clientSecret = viper.GetString("clientSecret")

	// OAuth2配置
	conf = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     basecamp.Endpoint,
		// 根据各网站的API的权限范围设置scopes
		// 不确定时，可以不设置
		Scopes: []string{""},
	}
}

var (
	oauthTypeWebServer = oauth2.SetAuthURLParam("type", "web_server")
	redirectURL        = "http://localhost:8080/callback" // 确保此重定向URL已在Basecamp应用程序设置中注册
)

func main() {

	// 创建一个HTTP服务器用于处理回调
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := conf.AuthCodeURL("", oauthTypeWebServer)
		http.Redirect(w, r, url, http.StatusFound)
	})

	//http.HandleFunc("/callback", newToken)
	http.HandleFunc("/callback", refreshToken)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

var (
	// 获取新的token
	newToken = func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		// 使用授权码交换获取token
		token, err := conf.Exchange(context.Background(), code, oauthTypeWebServer)
		if err != nil {
			http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		fmt.Fprintf(w, "<br><b>Token:</b> %s\n\n", token.AccessToken)
		fmt.Fprintf(w, "<br><b>Expiry</b>:</b> %s\n\n", token.Expiry)
		fmt.Fprintf(w, "<br><b>RefreshToken:</b> %s\n\n", token.RefreshToken)
	}

	// 刷新已有的token
	refreshToken = func(w http.ResponseWriter, r *http.Request) {
		conf.Endpoint = basecamp.Endpoint4Refresh

		existingToken := &oauth2.Token{
			//AccessToken:  "BAhbB0kiAbB7ImNsaWVudF9pZCI6IjM5ZTI0YjEwNzg1MDE4YTg3ZmVmNTg0YmNlZGE0MWIwMGRhZDQ5MDkiLCJleHBpcmVzX2F0IjoiMjAyMy0xMi0xM1QwNzo1OToyOFoiLCJ1c2VyX2lkcyI6WzQ0NjM4ODkyXSwidmVyc2lvbiI6MSwiYXBpX2RlYWRib2x0IjoiZGYzNTZmZTE3MmNjZmY5ZjMwMmExYzE2NmI0ZGFlYWMifQY6BkVUSXU6CVRpbWUNp+0ewMvtyu0JOg1uYW5vX251bWkCogE6DW5hbm9fZGVuaQY6DXN1Ym1pY3JvIgdBgDoJem9uZUkiCFVUQwY7AEY=--4f66aeeae635117d5293e297a62a1cf8a2f73d8e",
			RefreshToken: "BAhbB0kiAbB7ImNsaWVudF9pZCI6IjM5ZTI0YjEwNzg1MDE4YTg3ZmVmNTg0YmNlZGE0MWIwMGRhZDQ5MDkiLCJleHBpcmVzX2F0IjoiMjAzMy0xMS0yOVQwNzo1OToyOFoiLCJ1c2VyX2lkcyI6WzQ0NjM4ODkyXSwidmVyc2lvbiI6MSwiYXBpX2RlYWRib2x0IjoiZGYzNTZmZTE3MmNjZmY5ZjMwMmExYzE2NmI0ZGFlYWMifQY6BkVUSXU6CVRpbWUNp2shwAj6yu0JOg1uYW5vX251bWkCzwE6DW5hbm9fZGVuaQY6DXN1Ym1pY3JvIgdGMDoJem9uZUkiCFVUQwY7AEY=--b71d7a18d6011c00019386cc3ad2dc4f3e7a1c2d",
			//TokenType:    "Bearer",
			//Expiry:       time.Now().Add(-1 * time.Hour), // 设置为过去的时间以触发刷新
		}

		// 使用刷新令牌刷新令牌
		newToken, err := conf.TokenSource(context.Background(), existingToken).Token()
		if err != nil {
			http.Error(w, "Failed to refresh token", http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
		fmt.Fprintf(w, "<p>")
		fmt.Fprintf(w, "<br><b>Token:</b> %s\n\n", newToken.AccessToken)
		fmt.Fprintf(w, "<br><b>Expiry</b>:</b> %s\n\n", newToken.Expiry)
		fmt.Fprintf(w, "<br><b>RefreshToken:</b> %s\n\n", newToken.RefreshToken)
	}
)
