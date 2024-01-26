package impl

import (
	"context"
	"time"

	"github.com/kadegolang/sso/app/user"
	"gorm.io/gorm"
)

// 确保 UserServiceImpl 类型实现了 user.Service 接口。
// var _ user.Service = &UserServiceImpl{}
var _ user.Service = (*UserServiceImpl)(nil)

type UserServiceImpl struct {
	db *gorm.DB
}

// func NewUserServiceImpl() *UserServiceImpl {
// 	return &UserServiceImpl{
// 		// db: ioc.Controller().GetName(mysql.AppName).(mysql.Service),
// 		db: conf.C().MySQL.GetConn().Debug(),
// 		// 		dsn := "user:pass@tcp(35.220.136.110:3306)/sso?charset=utf8mb4&parseTime=True&loc=Local"
// 		//   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	}
// }

// create user
func (i *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.Users, error) {
	//1.验证账号密码是否为空
	if err := VerifyNameAndPassword(req); err != nil {
		return nil, err
	}
	//2.创建用户
	in := user.NewUsers(req)
	err := i.db.Model(&user.Users{}).Create(&in).Error
	return in, err
}

// query user
func (i *UserServiceImpl) QueryUser(ctx context.Context, req *user.QueryUserRequest) (*user.Users, error) {
	query := i.db.WithContext(ctx).Model(&user.Users{})
	if req.QueryByName != "" {
		query = query.Where("username = ?", req.QueryByName) //优先查询用户名
	} else {
		query = query.Where("id = ?", req.QueryByID)
	}
	in := user.NewUsers(&user.CreateUserRequest{})
	err := query.First(&in).Error
	return in, err
}

// edit user
func (i *UserServiceImpl) EditUser(ctx context.Context, req *user.EditUserRequest) (*user.Users, error) {
	in := user.NewUsers(&user.CreateUserRequest{})
	query := i.db.WithContext(ctx).Model(&user.Users{}).Where("username = ?", req.Username)
	err := query.Updates(map[string]interface{}{"update_time": time.Now().Unix(), "password": req.PassWord}).First(&in).Error
	return in, err
}

// delete user
func (i *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) error {
	query := i.db.WithContext(ctx).Model(&user.Users{})
	if req.Username != "" {
		query = query.Where("username = ?", req.Username) //优先查询用户名
	} else {
		query = query.Where("id = ?", req.Id)
	}
	err := query.Delete(&user.Users{}).Error
	return err
}
