package impl

import (
	"context"

	"gitee.com/go-kade/library/ioc"
	"gitee.com/go-kade/library/ioc/config/datasource"
)

const (
	AppName = "user"
)

// 导入这个包的时候，直接把这个对象 UserServiceImpL 注册给IocUserService ioc
// 注册user业务模块(业务模块的名称是user.AppName)的控制器
func init() {
	ioc.Default().Registry(&UserServiceImpl{})
}

func (u *UserServiceImpl) Init() error {
	u.db = datasource.DB(context.Background())
	return nil
}

func (u *UserServiceImpl) Name() string {
	return AppName
}

func (u *UserServiceImpl) Version() string {
	return ""
}

func (u *UserServiceImpl) Priority() int {
	return 9
}

func (u *UserServiceImpl) Close(ctx context.Context) error {
	return nil
}

func (u *UserServiceImpl) AllowOverwrite() bool {
	return true
}
