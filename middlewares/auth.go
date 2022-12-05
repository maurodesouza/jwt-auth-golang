package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(c *gin.Context) {
	authorization := c.Request.Header["Authorization"]

	if len(authorization) != 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token not provided",
		})

		return
	}

	splitedToken := strings.Split(authorization[0], " ")

	if len(splitedToken) != 2 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "badly formatted token",
		})

		return
	}

	tokenString := splitedToken[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token invalid",
		})

		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token invalid",
		})
	}

	var (
		sub = claims["sub"]
		exp = claims["exp"]
	)

	if float64(time.Now().Unix()) > exp.(float64) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token expired",
		})

		return
	}

	c.Set("user_id", sub)

	c.Next()
}
