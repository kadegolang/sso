package impl

import (
	"fmt"

	"github.com/kadegolang/sso/app/user"
)

// 验证用户名密码是否为空
func VerifyNameAndPassword(req *user.CreateUserRequest) error {
	if req.UserName == "" || req.PassWord == "" {
		return fmt.Errorf("username or password is cannot null")
	}
	return nil
}
