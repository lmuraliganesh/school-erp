package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims defines what data we store inside the token
type Claims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, role string) (string, error) {
	// 1. Build a Claims struct with UserID, Role, and
	//    RegisteredClaims{ ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)) }
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	// 2. token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3. Sign it: return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidateJWT(tokenString string) (*Claims, error) {
	// 1. claims := &Claims{}
	claims := &Claims{}

	// 2. _, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	//        return []byte(os.Getenv("JWT_SECRET")), nil
	//    })
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	// 3. if err != nil { return nil, err }
	if err != nil {
		return nil, err
	}

	// 4. return claims, nil
	return claims, nil
}
