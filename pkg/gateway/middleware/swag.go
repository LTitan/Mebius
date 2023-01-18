package middleware

import (
	_ "github.com/LTitan/Mebius/statik"
	"github.com/gin-gonic/gin"
	"github.com/rakyll/statik/fs"
)

func SwaggerDoc(server *gin.Engine) error {
	handle, err := fs.New()
	if err != nil {
		return err
	}
	server.StaticFS("/swagger", handle)
	return nil
}
