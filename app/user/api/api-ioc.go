package api

import (
	"context"

	"gitee.com/go-kade/library/ioc"
	"github.com/kadegolang/sso/app/user"
	"github.com/kadegolang/sso/app/user/impl"
)

// 导入这个包的时候，直接把这个对象 UserServiceImpL 注册给IocUserService ioc
// 注册user业务模块(业务模块的名称是user.AppName)的控制器
func init() {
	ioc.Api().Registry(&UserApiHeadler{})
}

type UserApiHeadler struct {
	user user.Service //依赖控制器
	// user *impl.UserServiceImpl //依赖控制器 不应该这样干他的实现，应该user.Service方法才对

}

func (u *UserApiHeadler) Init() error {
	u.user = ioc.Default().Get(impl.AppName).(*impl.UserServiceImpl)
	return nil
}

func (u *UserApiHeadler) Name() string {
	return "users"
}

func (u *UserApiHeadler) Version() string {
	return ""
}

func (u *UserApiHeadler) Priority() int {
	return 9
}

func (u *UserApiHeadler) Close(ctx context.Context) error {
	return nil
}

func (u *UserApiHeadler) AllowOverwrite() bool {
	return true
}
