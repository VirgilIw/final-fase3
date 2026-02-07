package pkg

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/virgilIw/final-fase3/internal/dto"
)

type JwtClaims struct {
	*dto.JWTClaims
}

func NewJWTClaims(id int, role string) *JwtClaims {
	return &JwtClaims{
		JWTClaims: &dto.JWTClaims{
			UserID: id,
			Role:   role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
				Issuer:    os.Getenv("JWT_ISSUER"),
			},
		},
	}
}

func (jc *JwtClaims) GenToken() (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("jwt secret not found")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jc)
	return token.SignedString([]byte(jwtSecret))
}

func (jc *JwtClaims) VerifyToken(token string) (bool, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return false, errors.New("jwt secret not found")
	}

	jwtToken, err := jwt.ParseWithClaims(token, jc, func(t *jwt.Token) (any, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return false, errors.New("token expired")
		}
		return false, err
	}

	if !jwtToken.Valid {
		return false, errors.New("invalid token")
	}

	iss, err := jwtToken.Claims.GetIssuer()
	if err != nil {
		return false, errors.New("invalid token claims")
	}

	expectedIssuer := os.Getenv("JWT_ISSUER")
	if expectedIssuer == "" {
		return false, errors.New("issuer not found")
	}

	if iss != expectedIssuer {
		return false, errors.New("invalid issuer")
	}

	return true, nil
}
