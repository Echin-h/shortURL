package url

import (
	"ShortUrl/internal/module/url/function"
	"github.com/gin-gonic/gin"
)

func (u *ModuleURL) InitRouter(r *gin.RouterGroup) {
	r.POST("/shorten", function.Shorten)
	r.GET("/:short", function.Longer)
	r.GET("", function.Ping)
}
