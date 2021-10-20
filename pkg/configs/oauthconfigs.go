package configs

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

// InitGoogleOAuthConfigs read configs from envs\config files and returns *oauth2.Config for Google
func InitGoogleOAuthConfigs() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  "http://localhost:8080/oauth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}
