package middleware

import (
	"errors"
	"strings"
	"time"

	"github.com/DimasPramantya/goMiniProject/utils/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type Claims struct {
	jwt.StandardClaims
}

var secret = viper.GetString("SECRET_KEY")

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := GetJwtTokenFromHeader(c)
		if err != nil {
			helper.WriteError(c, 401, "Unauthorized", nil)
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unauthorized") 
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			helper.WriteError(c, 401, "Invalid token", nil)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			helper.WriteError(c, 401, "Invalid token claims", nil)
			c.Abort()
			return
		}

		// Set auth context data
		c.Set("user_id", claims["user_id"])
		c.Set("username", claims["username"])

		c.Next()
	}
}

func GetJwtTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	// Expecting format: Bearer <token>
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("authorization header format must be Bearer {token}")
	}

	return parts[1], nil
}

//TOKEN EXPIRED IN 1 DAY
func GenerateJwtToken(username string, userId int) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userId,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), 
		"iat":      time.Now().Unix(),                    
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}