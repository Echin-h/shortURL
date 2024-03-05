package url

import "github.com/gin-gonic/gin"

func (u *ModuleURL) InitRouter(r *gin.RouterGroup) {
	r.POST("/shorten", Shorten)
}
