package module

import (
	"ShortUrl/internal/module/url"
	"github.com/gin-gonic/gin"
)

type Module interface {
	GetName() string
	Init()
	InitRouter(r *gin.RouterGroup)
}

var Modules []Module

func RegisterRouters(m Module) {
	Modules = append(Modules, m)
}

func init() {
	RegisterRouters(&url.ModuleURL{})
}
