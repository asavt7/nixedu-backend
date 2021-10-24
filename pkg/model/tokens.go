package model

// CachedTokens -  access tokens cached to auth storage
type CachedTokens struct {
	AccessUID  string `json:"access"`
	RefreshUID string `json:"refresh"`
}
