package middleware

import (
	"go_chat/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		userClaims, err := helper.AnalyseToken(token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "User authentication fails. Procedure",
			})
			return
		}
		c.Set("user_claims", userClaims)
		c.Next()
	}
}
