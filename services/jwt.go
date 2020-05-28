package services

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))
var anotherJwtKey = []byte(os.Getenv("ANOTHER_SECRET_KEY"))
var refreshTokenKey = []byte(os.Getenv("REFRESH_TOKEN_KEY"))

// Claims defines jwt claims
type Claims struct {
	UserID string `json:"email"`
	jwt.StandardClaims
}

// GenerateToken handles generation of a jwt code
// @returns string -> token and error -> err
func GenerateToken(userID string) (string, string, error) {
	// Define token expiration time
	expirationTime := time.Now().Add(1440 * time.Minute)
	// Refresh token expiration time (3 days)
	refreshExpTime := time.Now().Add(4320 * time.Minute)
	// Define the payload and exp time
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	refreshClaims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpTime.Unix(),
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	// Sign token with secret key encoding
	tokenString, err := token.SignedString(jwtKey)
	refreshTokenString, err := refreshToken.SignedString(refreshTokenKey)

	return tokenString, refreshTokenString, err
}

// GenerateNonAuthToken handles generation of a jwt code
// @returns string -> token and error -> err
func GenerateNonAuthToken(userID string) (string, error) {
	// Define token expiration time
	expirationTime := time.Now().Add(1440 * time.Minute)
	// Define the payload and exp time
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key encoding
	tokenString, err := token.SignedString(anotherJwtKey)

	return tokenString, err
}

// DecodeNonAuthToken handles decoding a jwt token
func DecodeNonAuthToken(tkStr string) (string, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tkStr, claims, func(token *jwt.Token) (interface{}, error) {
		return anotherJwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}

	if !tkn.Valid {
		return "", err
	}

	return claims.UserID, nil
}

// DecodeToken handles decoding a jwt token
func DecodeToken(tkStr string) (string, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tkStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}

	if !tkn.Valid {
		return "", err
	}

	return claims.UserID, nil
}

// DecodeRefreshToken handles decoding refresh token
func DecodeRefreshToken(tkStr string) (string, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tkStr, claims, func(token *jwt.Token) (interface{}, error) {
		return refreshTokenKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
		return "", err
	}

	if !tkn.Valid {
		return "", err
	}

	return claims.UserID, nil
}
