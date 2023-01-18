package middleware

import (
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		klog.Infof("[%s] from client: %s, request: %s\n", c.Request.Method, c.ClientIP(), c.Request.URL)
		c.Next()
	}
}
