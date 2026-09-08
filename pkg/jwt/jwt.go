package jwt

import (
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	jwtlib.RegisteredClaims
}

type Service struct {
	secret []byte
}

func NewService(secret string) *Service {
	return &Service{secret: []byte(secret)}
}

var (
	ErrTokenExpired = jwtlib.ErrTokenExpired
)

func (s *Service) GenerateToken(userID uuid.UUID, ttl time.Duration) (string, error) {
	now := time.Now()
	expiredAt := now.Add(ttl)
	claims := Claims{
		RegisteredClaims: jwtlib.RegisteredClaims{
			Subject:   userID.String(),
			ExpiresAt: jwtlib.NewNumericDate(expiredAt),
			IssuedAt:  jwtlib.NewNumericDate(now),
		},
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}

func (s *Service) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwtlib.ParseWithClaims(tokenString, &Claims{}, func(token *jwtlib.Token) (any, error) {
		if _, ok := token.Method.(*jwtlib.SigningMethodHMAC); !ok {
			return nil, jwtlib.ErrSignatureInvalid
		}
		return s.secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwtlib.ErrSignatureInvalid
	}

	return claims, nil
}
