package impl

import (
	"context"
	"fmt"
	"time"

	"gitee.com/go-kade/library/exception"
	"github.com/kadegolang/sso/app/token"
	"github.com/kadegolang/sso/app/user"
	"github.com/kadegolang/sso/tools"
	"github.com/rs/xid"
	"gorm.io/gorm"
)

// 接口约束，需要满足以下方法
// type Service interface {
// 	Login(context.Context, *LoginRequst) (*Token, error)         //登灵接口（额发Token)
// 	Logout(context.Context, *LoginOutRequst) error               //登灵接口（销毁Token)
// 	ValiateToken(context.Context, *ValiateToken) (*Token, error) //校验Token 是给內部中间层使用 身份校验层
// }

var _ token.Service = (*TokenServiceImpl)(nil)

type TokenServiceImpl struct {
	db *gorm.DB
	//这里需要依赖另一个业务领城：用户管理领城
	user user.Service
}

// 登灵接口（额发Token)
func (i *TokenServiceImpl) Login(ctx context.Context, req *token.LoginRequst) (*token.Token, error) {
	//1.查询用户
	ureq := user.NewQueryUserRequest()
	ureq.QueryByName = req.Username
	u, err := i.user.QueryUser(ctx, ureq)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// fmt.Println(exception.NewAuthFound("cccccccc").Error())
			return nil, exception.NewNotFound("record not found,用户%s查询失败", req.Username)
		}
		return nil, err
	}

	//2.比对密码
	err = tools.CheckPassword(u.Password, req.Password) //err = u.CheckPassword(req.Password)
	if err != nil {
		return nil, exception.NewNotFound("用户名或者密码不正确:%s", err)
	}
	//3.颁发token
	tk := token.NewToken()
	tk.UserID = u.Id
	tk.UserName = u.Username
	switch i.db.WithContext(ctx).Where("username = ?", u.Username).First(&tk).Error {
	// 避免同一个用户多次登灵
	case gorm.ErrRecordNotFound: //3.颁发token,保存token
		i.db.WithContext(ctx).Create(&tk)
		return tk, nil
	case nil: //4.更新token
		tk.AccessToken = tk.RefreshToken
		tk.UpdatedAt = time.Now().Unix()
		// tk.CreatedAt = time.Now().Unix()
		tk.RefreshToken = xid.New().String()
		i.db.WithContext(ctx).Where("username= ?", tk.UserName).Updates(tk)
		return tk, nil
	default:
		return tk, exception.NewNotFound("错误异常")
	}
}

// 登灵接口（销毁Token)
func (i *TokenServiceImpl) Logout(ctx context.Context, req *token.LoginOutRequst) error {
	// tk := token.NewToken()
	err := i.db.Where("access_token = ?", req.AccessToken).Delete(&token.Token{}).Error
	if err != nil {
		return exception.NewNotFound("token check failure")
	}
	return nil
}

// 校验Token 是给內部中间层使用 身份校验层
func (c *TokenServiceImpl) ValiateToken(ctx context.Context, req *token.ValiateTokenRequst) (*token.Token, error) {
	//1.查询token（谁颁发的）
	fmt.Println("access_token", req.AccessToken)
	tk := token.NewToken() //结构体先初始化，不然查询数据库会报空指针
	err := c.db.WithContext(ctx).Where("access_token = ?", req.AccessToken).First(&tk).Error
	// fmt.Println("-------1177171", tk.UserID, tk.UserName)
	if err != nil {
		return nil, exception.NewNotFound("token check failure")
	}
	//2.判断token的合法性
	access_token_expired_at_seconds, err := tk.IsExpired()
	if err != nil {
		return nil, err
	}
	// fmt.Println("-------1177173", tk.UserID, tk.UserName, access_token_expired_at_seconds)
	// 补充用户信息, 只补充了用户的角色
	in := user.NewQueryUserRequest()
	in.QueryByID = tk.UserID
	in.QueryByName = tk.UserName
	//3.查询用户
	u, err := c.user.QueryUser(ctx, in)
	if err != nil {
		return nil, err
	}
	//4.把查询到的用户返回给tk
	tk.Role = u.Role
	tk.Access_token_expired_at_seconds = access_token_expired_at_seconds
	return tk, err
}
