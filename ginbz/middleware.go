package ginbz

import (
	"log"
	"net/http"

	"github.com/bigzhu/gobz/apibz"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware 必须要求登录
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user")
		if userID == nil {
			log.Println("get userID", c.MustGet("userID"))
			c.AbortWithStatusJSON(http.StatusUnauthorized, apibz.NewE("need login"))
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}
