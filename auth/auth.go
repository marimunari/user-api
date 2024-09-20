package auth

import (
	"errors"
	"os"
	"sync"
	"time"
	"user-api/models"

	"github.com/golang-jwt/jwt/v5"
)

var (
	JWTKey         = []byte(os.Getenv("JWT_SECRET"))
	RefreshJWTKey  = []byte(os.Getenv("REFRESH_JWT_SECRET"))
	blacklist      = make(map[string]time.Time)
	blacklistMutex sync.Mutex
)

// Method to generate a JWT token
func GenerateToken(email string) (string, string, error) {
	accessExpirationTime := time.Now().Add(1 * time.Hour)
	refreshExpirationTime := time.Now().Add(24 * time.Hour)

	claims := &models.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpirationTime),
			Issuer:    "user-api",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString(JWTKey)
	if err != nil {
		return "", "", err
	}

	refreshClaims := &models.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpirationTime),
			Issuer:    "user-api",
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(RefreshJWTKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

// Method to check JWT token validity
func ValidateToken(tokenString string, isRefreshToken bool) (*models.Claims, error) {
	claims := &models.Claims{}

	var token *jwt.Token
	var err error

	if isRefreshToken {
		token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("método de assinatura inválido")
			}
			return RefreshJWTKey, nil
		})
	} else {
		token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("método de assinatura inválido")
			}
			return JWTKey, nil
		})
	}

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}

// Method to check if the token is on the blacklist
func IsTokenBlacklisted(token string) bool {
	blacklistMutex.Lock()
	defer blacklistMutex.Unlock()

	_, exists := blacklist[token]
	return exists
}

// Method to add a token to the blacklist with an expiration time
func AddTokenToBlacklist(token string) {
	blacklistMutex.Lock()
	defer blacklistMutex.Unlock()

	blacklist[token] = time.Now().Add(time.Hour * 24)
}
