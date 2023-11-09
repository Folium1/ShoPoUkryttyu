package models

type JWT struct {
	CookieKey string `json:"cookie_key"`
	Token     string `json:"token"`
}
