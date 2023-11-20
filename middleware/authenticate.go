package middleware

import (
	"ki_assignment-1/service"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		if !strings.Contains(authHeader, "Bearer") {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		userID, err := jwtService.FindUserIDByToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set("token", tokenString)
		c.Set("userID", userID)
		c.Next()
	}
}
