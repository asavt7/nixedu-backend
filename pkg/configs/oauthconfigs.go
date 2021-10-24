package configs

import (
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// InitGoogleOAuthConfigs read configs from envs\config files and returns *oauth2.Config for Google
func InitGoogleOAuthConfigs() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  "http://localhost:8080/oauth/google/callback",
		ClientID:     viper.GetString("oauth.google.client.id"),
		ClientSecret: viper.GetString("oauth.google.client.secret"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}
