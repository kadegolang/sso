package token

import "context"

type Service interface {
	Login(context.Context, *LoginRequst) (*Token, error)               //登灵接口（额发Token)
	Logout(context.Context, *LoginOutRequst) error                     //登灵接口（销毁Token)
	ValiateToken(context.Context, *ValiateTokenRequst) (*Token, error) //校验Token 是给內部中间层使用 身份校验层
}

type LoginRequst struct {
	// Username string `json:"username"`
	// Password string `json:"password"`

	Username string `form:"username"`
	Password string `form:"password"`
}

// Name     string `form:"name"`
// Password string `form:"password"`

// 让api调用使用的
func NewLoginRequst() *LoginRequst {
	return &LoginRequst{}
}

// 万一的Token滋露，不知道refresh_token，也没法推出
type LoginOutRequst struct {
	AccessToken string `json:"access_token"` //颁发给用户的访问令牌（需要用户携带token来访问接口）
	// RefreshToken string `json:"refresh_token"` //刷新令牌
}

// 让api调用使用的
func NewLoginOutRequst() *LoginOutRequst {
	return &LoginOutRequst{}
}

type ValiateTokenRequst struct {
	AccessToken string `json:"access_token"`
}

// 让api调用使用的
func NewValiateTokenRequst(access_token_web string) *ValiateTokenRequst {
	return &ValiateTokenRequst{
		AccessToken: access_token_web,
	}
}
