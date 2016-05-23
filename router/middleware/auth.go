package middleware

import (
	// "fmt"
	"time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/yhsiang/bass/store"
)

func Auth() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour * 24,
		Authenticator: func(userId string, password string, c *gin.Context) (string, bool) {
			user, err := store.GetUserBy(c, map[string]interface{}{
				"username": userId,
			})
			if err != nil {
				return userId, false
			}

			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err != nil {
				return userId, false
			}
			return userId, true
		},
		Authorizator: func(userId string, c *gin.Context) bool {
			_, err := store.GetUser(c, userId)
			if err != nil {
				return false
			}
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	}
}
