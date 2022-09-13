package middleware

import (
	"musicAPI/src/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "no token authorized",
			})
			c.Abort()
			return
		}
		claims, err := helper.ValidateToken(clientToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}
		c.Set("userID", claims)

		c.Next()

	}
}
