package jwt

import (
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTProvider struct {
	secret []byte
}

func NewJWTProvider(secret string) *JWTProvider {
	return &JWTProvider{
		secret: []byte(secret),
	}
}

func (p *JWTProvider) Generate(userID uuid.UUID, ttl time.Duration) (string, error) {
	now := time.Now()
	expiresAt := now.Add(ttl)
	claims := jwtlib.MapClaims{
		"sub": userID.String(),
		"iat": now,
		"exp": expiresAt,
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)

	return token.SignedString(p.secret)
}
