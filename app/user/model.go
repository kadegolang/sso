package user

import (
	"encoding/json"
	"time"

	"github.com/kadegolang/sso/common"
)

type Users struct {
	*common.Meta
	Username string `json:"username" gorm:"column:username;comment:用户名, 用户名不允许重复的"`
	Password string `json:"password" gorm:"column:password;comment:不能保持用户的明文密码,hash算法"`
	Role     Roles  `json:"role" gorm:"column:role;comment:用户角色,0表示普通用户,1表示审核用户,2表示管理员"`
}

// 创建用户，需要创建对象
func NewUsers(req *CreateUserRequest) *Users {
	//1.判断是否传入进来标签，如果没有空值
	if req.Label == nil {
		req.Label = map[string]string{}
	}
	//2.密码需要hash加密
	req.HashPassword()
	return &Users{
		Meta: &common.Meta{
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
			Label:      req.Label,
		},
		Username: req.UserName,
		Password: req.PassWord,
		Role:     req.Role,
	}
}

// 返回表名
func (*Users) TableName() string {
	return "users"
}

func (u *Users) String() string {
	u1, _ := json.Marshal(u)
	return string(u1)
}

/*
//检查密码,暂时可以不用
func (u *Users) CheckPassword(password string) error {
	// fmt.Println("password", password)
	// fmt.Println("password", string(password))
	// fmt.Println("u.Password", string(u.Password))
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
*/
