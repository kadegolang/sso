package impl_test

import (
	"context"
	"testing"

	"gitee.com/go-kade/library/ioc"
	_ "gitee.com/go-kade/library/ioc/config/datasource"
	"github.com/kadegolang/sso/app/user"
	_ "github.com/kadegolang/sso/app/user/api"
	"github.com/kadegolang/sso/app/user/impl"
	"github.com/gin-gonic/gin"
)

func TestCreateUser(t *testing.T) {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go12-project/sso/etc/application.toml"
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
	ioc.LoadGinApi("/api/v1", r)

	// 启动 Gin 服务器
	r.Run(":8080")

}

// query user
func TestQueryUser(t *testing.T) {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go12-project/sso/etc/application.toml"
	ioc.DevelopmentSetup(req)
	a := ioc.Default().Get(impl.AppName).(*impl.UserServiceImpl)
	// req := user.NewCreateUserRequest()
	// req.UserName = "sja"
	// req.PassWord = "sman"
	// req.Label = map[string]string{"bsab": "111111"}

	t.Log(a.QueryUser(context.Background(), &user.QueryUserRequest{QueryByName: "22q32112", QueryByID: 31111}))
}

// edit user
func TestEditUser(t *testing.T) {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go12-project/sso/etc/application.toml"
	ioc.DevelopmentSetup(req)
	a := ioc.Default().Get(impl.AppName).(*impl.UserServiceImpl)
	// req := user.NewCreateUserRequest()
	// req.UserName = "sja"
	// req.PassWord = "sman"
	// req.Label = map[string]string{"bsab": "111111"}
	b, err := a.EditUser(context.Background(), &user.EditUserRequest{Username: "2222", PassWord: "cys01100522"})
	t.Log(b, err)
}

// delete user
func TestDeleteUser(t *testing.T) {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go12-project/sso/etc/application.toml"
	ioc.DevelopmentSetup(req)
	a := ioc.Default().Get(impl.AppName).(*impl.UserServiceImpl)
	// req := user.NewCreateUserRequest()
	// req.UserName = "sja"
	// req.PassWord = "sman"
	// req.Label = map[string]string{"bsab": "111111"}
	err := a.DeleteUser(context.Background(), &user.DeleteUserRequest{Id: 32})
	t.Log(err)
}
