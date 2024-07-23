package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Get("auth"))
		c.Next()
	}
}
