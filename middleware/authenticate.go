package middleware

import (
	"learn-gin/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authen(c *gin.Context) {
	token := c.Request.Header.Get("x-access-token")
	userID, err := util.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Set("user", userID)
	c.Next()
}
