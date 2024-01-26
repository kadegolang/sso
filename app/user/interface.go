package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(context.Context, *CreateUserRequest) (*Users, error)
	QueryUser(context.Context, *QueryUserRequest) (*Users, error)
	EditUser(context.Context, *EditUserRequest) (*Users, error)
	DeleteUser(context.Context, *DeleteUserRequest) error
}

type CreateUserRequest struct {
	UserName string            `json:"username" gorm:"column:username" comment:"用户名, 用户名不允许重复的"`
	PassWord string            `json:"password" gorm:"column:password" comment:"不能保持用户的明文密码,hash算法"`
	Label    map[string]string `json:"label" gorm:"column:label;serializer:json"`
	Role     Roles             `json:"role" gorm:"column:role;comment:用户角色,0表示普通用户,1表示审核用户,2表示管理员"`
}

// 暂时可无
func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Label: map[string]string{},
	}
}

// 密码hash加密
func (req *CreateUserRequest) HashPassword() {
	pass, _ := bcrypt.GenerateFromPassword([]byte(req.PassWord), bcrypt.DefaultCost)
	req.PassWord = string(pass)
}

type QueryUserRequest struct {
	QueryByID int64 `json:"id"` //当queryby等于0时，传进去的就是id，1时就是nanme
	// QueryBy       QueryUserBy //当queryby等于0时，传进去的就是id，1时就是nanme
	QueryByName string `json:"username" gorm:"column:username;comment:用户名, 用户名不允许重复的"`
}

// 创建QueryUserRequest结构体
func NewQueryUserRequest() *QueryUserRequest {
	return &QueryUserRequest{}
}

type EditUserRequest struct {
	Username string `json:"username" gorm:"column:username;comment:用户名, 用户名不允许重复的"`
	PassWord string `json:"password" gorm:"column:password;comment:不能保持用户的明文密码,hash算法"`
}

// 创建EditUserRequest结构体  此处给api调用时使用的
func NewEditUserRequest() *EditUserRequest {
	return &EditUserRequest{}
}

type DeleteUserRequest struct {
	Id       int64  `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username;comment:用户名, 用户名不允许重复的"`
}

// 创建DeleteUserRequest结构体  此处给api调用时使用的
func NewDeleteUserRequest() *DeleteUserRequest {
	return &DeleteUserRequest{}
}
