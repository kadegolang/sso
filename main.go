package main

import (
	"gitee.com/go-kade/library/ioc"
	// _ "gitee.com/go-kade/library/ioc/config/datasource"  初始化时做过了
	_ "github.com/kadegolang/sso/app/token/api"
	_ "github.com/kadegolang/sso/app/user/api"
	"github.com/gin-gonic/gin"
)

func main() {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "etc/application.toml"
	ioc.DevelopmentSetup(req)
	// a := ioc.Default().Get(impl.AppName).(*impl.UserServiceImpl)
	// req1 := user.NewCreateUserRequest()
	// req1.UserName = "sja11"
	// req1.PassWord = "sman11"
	// req1.Label = map[string]string{"bsa1b": "111111"}
	// t.Log(a.CreateUser(context.Background(), req1))
	// 加载 Gin API，指定路径前缀和路由器对象
	r := gin.Default()
	// // a := r.Group("ccc")
	// r.POST("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "Hello from Gin API"})
	// })
	r.LoadHTMLGlob("views/*") //配置模版中间件
	ioc.LoadGinApi("/api/v1", r)
	// 启动 Gin 服务器
	r.Run(":8080")
}
