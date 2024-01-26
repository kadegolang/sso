package impl_test

import (
	"context"
	"fmt"
	"testing"

	"gitee.com/go-kade/library/ioc"
	"github.com/kadegolang/sso/app/token"
	"github.com/kadegolang/sso/app/token/impl"
	"github.com/kadegolang/sso/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "/Users/kade.chen/go12-project/sso/etc/application.toml"
	ioc.DevelopmentSetup(req)
}

func TestLogin(t *testing.T) {
	a := ioc.Default().Get(impl.AppName).(*impl.TokenServiceImpl)
	fmt.Println(ioc.Default().List(), ioc.Config().List(), ioc.Controller().List())
	// impl.NewTokenServiceImpl(&userimpl.UserServiceImpl{}).Login(context.Background(), &token.LoginRequst{Username: "a1111111", Password: "a1111"})
	tk, err := a.Login(context.Background(), &token.LoginRequst{Username: "cc", Password: "123456"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func TestLoginOut(t *testing.T) {
	a := ioc.Default().Get(impl.AppName).(*impl.TokenServiceImpl)
	fmt.Println(ioc.Default().List(), ioc.Config().List(), ioc.Controller().List())
	// impl.NewTokenServiceImpl(&userimpl.UserServiceImpl{}).Login(context.Background(), &token.LoginRequst{Username: "a1111111", Password: "a1111"})
	err := a.Logout(context.Background(), &token.LoginOutRequst{AccessToken: "cmadedldrb6lrmk03c70"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestValiateToken(t *testing.T) {
	a := ioc.Default().Get(impl.AppName).(*impl.TokenServiceImpl)
	fmt.Println(ioc.Default().List(), ioc.Config().List(), ioc.Controller().List())
	// impl.NewTokenServiceImpl(&userimpl.UserServiceImpl{}).Login(context.Background(), &token.LoginRequst{Username: "a1111111", Password: "a1111"})
	in, err := a.ValiateToken(context.Background(), token.NewValiateTokenRequst("cmifr4ddrb6o3sts94l0"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(in)
}


func TestValiateToken1(t *testing.T) {
	middlewares.NewTokenAuther().Auth(&gin.Context{})
}
