package middleware

import (
	"strings"

	"github.com/Raghunandan-79/auth-service/utils"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "missing token"})
			return 
		}

		token := strings.TrimPrefix(header, "Bearer ")
		claims, err := utils.ValidateAccessToken(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
			return 
		}

		c.Set("user_id", claims.UserID)
		c.Next()
	}
}