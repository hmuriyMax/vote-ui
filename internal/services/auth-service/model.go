package auth_service

import "time"

const (
	CookieName   = "auth_token"
	CookieUserId = "user_id"
)

type User struct {
	ID    int64  `json:"id"`
	Token *Token `json:"token"`
}

type Token struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires"`
}
