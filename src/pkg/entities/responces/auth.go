package responces

import "github.com/golang-jwt/jwt/v4"

type AuthResponse struct {
	AccessToken  string
	RefreshToken string
}

type Claims struct {
	UserID string `json:"sub"`
	Exp    int64  `json:"exp"`
	jwt.RegisteredClaims
}
