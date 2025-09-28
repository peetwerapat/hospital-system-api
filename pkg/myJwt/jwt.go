package myJwt

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/peetwerapat/hospital-system-api/internal/domain"
)

func GetJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "default_secret"
	}
	return secret
}

func GetRefreshSecret() string {
	secret := os.Getenv("REFRESH_TOKEN_SECRET")
	if secret == "" {
		secret = "default_secret"
	}
	return secret
}

type Claims struct {
	ID         int    `json:"id"`
	HospitalID string `json:"hospitalId"`
	jwt.RegisteredClaims
}

func CreateToken(staff *domain.Staff, expiration time.Duration, isAccess bool) (string, error) {
	claims := &Claims{
		ID:         staff.ID,
		HospitalID: strconv.Itoa(staff.HospitalID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := GetRefreshSecret()
	if isAccess {
		secret = GetJWTSecret()
	}

	return token.SignedString([]byte(secret))
}

func VerifyRefreshToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(GetRefreshSecret()), nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
