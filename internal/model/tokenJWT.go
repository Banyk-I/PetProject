package model

import (
	"errors"
	"time"
)

type TokenClaims struct {
	ExpiresAt int64  `json:"exp"` // Час закінчення дії токена
	IssuedAt  int64  `json:"iat"` // Час видачі токена
	UserID    string `json:"user_id"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func (t *TokenClaims) Valid() error {
	if t.ExpiresAt < time.Now().Unix() {
		return errors.New("token is expired")
	}
	return nil
}
