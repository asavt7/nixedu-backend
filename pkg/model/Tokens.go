package model

type CachedTokens struct {
	AccessUID  string `json:"access"`
	RefreshUID string `json:"refresh"`
}
