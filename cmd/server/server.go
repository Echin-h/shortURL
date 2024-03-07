package server

import (
	"ShortUrl/configs"
	"ShortUrl/internal/global/database"
	"ShortUrl/internal/global/log"
	"ShortUrl/internal/module"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func Init() {
	log.Init() // 未测试，本来想输出到命令行中的，明天再研究一下
	configs.Init()
	database.Init()
	color.Red("it is red color")

	for _, m := range module.Modules {
		m.Init()
	}
}

func Run() {
	r := gin.Default()
	for _, m := range module.Modules {
		m.InitRouter(r.Group("/" + m.GetName()))
	}
	panic(r.Run(":8080"))
}
