package api

import (
	"context"

	"gitee.com/go-kade/library/ioc"
	"github.com/kadegolang/sso/app/token"
	tokenimpl "github.com/kadegolang/sso/app/token/impl"
)

// 导入这个包的时候，直接把这个对象 TokenApiHeadler TokenApiHeadler ioc
// 注册user业务模块(业务模块的名称是user.AppName)的控制器
func init() {
	ioc.Api().Registry(&TokenApiHeadler{})
}

type TokenApiHeadler struct {
	//因为这个api 依赖他的具体实现
	token token.Service
	// user *impl.UserServiceImpl //依赖控制器 不应该这样干他的实现，应该user.Service方法才对

}

func (u *TokenApiHeadler) Init() error {
	u.token = ioc.Default().Get(tokenimpl.AppName).(*tokenimpl.TokenServiceImpl)
	return nil
}

func (u *TokenApiHeadler) Name() string {
	return "tokens"
}

func (u *TokenApiHeadler) Version() string {
	return ""
}

func (u *TokenApiHeadler) Priority() int {
	return 9
}

func (u *TokenApiHeadler) Close(ctx context.Context) error {
	return nil
}

func (u *TokenApiHeadler) AllowOverwrite() bool {
	return true
}
