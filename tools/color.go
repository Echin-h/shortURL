package tools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
)

// 可以研究一下gin的颜色输出表示
// 优雅的关机
// 配置SSH证书登录

func ColorPrint() {
	str := color.Red.Sprintf("This is red")
	fmt.Println(str)
	gin.ForceConsoleColor()

}

/*
自定义域名
一些短链服务提供商允许用户使用自己的域名来创建短链。这通常涉及到以下几个步骤：

拥有一个域名：用户需要拥有一个域名，无论是购买的还是通过其他方式获得的。
进行DNS配置：用户需要在域名管理后台添加相应的DNS记录，例如CNAME记录，将域名指向短链服务提供商的服务器。
绑定域名：在短链服务提供商的管理后台，将自定义的域名与账户关联起来。
生成短链：在管理后台生成短链接地址，并选择自己的域名。
例如，如果用户有一个域名example.com，并且希望在example.com/short路径下生成短链，那么用户就需要在域名管理后台添加一个CNAME记录，将short指向短链服务提供商的服务器。然后在短链服务提供商的后台中，将example.com/short设置为自定义域名，并生成短链12。

*/
