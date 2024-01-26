package impl

import (
	"context"

	"gitee.com/go-kade/library/ioc"
	"gitee.com/go-kade/library/ioc/config/datasource"
	"github.com/kadegolang/sso/app/user/impl"
)

var (
	AppName = "token"
)

// 导入这个包的时候，直接把这个对象 UserServiceImpL 注册给IocUserService ioc
// 注册user业务模块(业务模块的名称是user.AppName)的控制器，想要注册需要满足，下面的方法
func init() {
	ioc.Default().Registry(&TokenServiceImpl{})
}

func (u *TokenServiceImpl) Init() error {
	u.db = datasource.DB(context.Background())
	u.user = ioc.Default().Get(impl.AppName).(*impl.UserServiceImpl)
	return nil
}

func (u *TokenServiceImpl) Name() string {
	return AppName
}

func (u *TokenServiceImpl) Version() string {
	return ""
}

func (u *TokenServiceImpl) Priority() int {
	return 9
}

func (u *TokenServiceImpl) Close(ctx context.Context) error {
	return nil
}

func (u *TokenServiceImpl) AllowOverwrite() bool {
	return true
}
