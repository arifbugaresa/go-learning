package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-learning/utils/common"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		trace_id := common.GenerateRandomString(6)
		c.Set("trace_id", trace_id)

		c.Next()
	}
}
