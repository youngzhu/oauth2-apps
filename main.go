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

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

var (
	oauthTypeWebServer = oauth2.SetAuthURLParam("type", "web_server")
	redirectURL        = "http://localhost:8080/callback" // 确保此重定向URL已在Basecamp应用程序设置中注册
)

func main() {
	clientID := viper.GetString("clientID")
	clientSecret := viper.GetString("clientSecret")

	// OAuth2配置
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     basecamp.Endpoint,
		// 根据各网站的API的权限范围设置scopes
		// 不确定时，可以不设置
		Scopes: []string{""},
	}

	// 创建一个HTTP服务器用于处理回调
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := conf.AuthCodeURL("", oauthTypeWebServer)
		http.Redirect(w, r, url, http.StatusFound)
	})

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		token, err := conf.Exchange(context.Background(), code, oauthTypeWebServer)
		if err != nil {
			http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
			log.Fatal(err)
			return
		}

		fmt.Fprintf(w, "<br><b>Token:</b> %s\n\n", token.AccessToken)
		fmt.Fprintf(w, "<br><b>Expiry</b>:</b> %s\n\n", token.Expiry)
		fmt.Fprintf(w, "<br><b>RefreshToken:</b> %s\n\n", token.RefreshToken)
	})

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
