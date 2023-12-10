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
			c.AbortWithStatusJSON(401, gin.H{"error": "Token Tidak Ditemukan"})
			return
		}
		if !strings.Contains(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token Tidak Valid"})
			return
		}
		authHeader = strings.ReplaceAll(authHeader, "Bearer ", "")
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token Tidak Valid"})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Akses Ditolak"})
			return
		}

		userID, err := jwtService.FindUserIDByToken(authHeader)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		c.Set("token", authHeader)
		c.Set("userID", userID)
		c.Next()
	}
}
