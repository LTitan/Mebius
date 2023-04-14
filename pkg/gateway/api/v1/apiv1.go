package v1

import "github.com/gin-gonic/gin"

func RegisterRouter(apiv1 *gin.RouterGroup) {
	apiv1.GET("/test", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
}
