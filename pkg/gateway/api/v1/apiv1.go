package v1

import "github.com/gin-gonic/gin"

func RegisterRouter(apiv1 *gin.RouterGroup) {
	apiv1.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"code": 0, "data": "ok"})
	})
}
