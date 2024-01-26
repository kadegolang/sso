package token

import (
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/kadegolang/sso/app/user"
	"github.com/kadegolang/sso/exception"
	"github.com/rs/xid"
)

type Token struct {
	UserID                int64  `json:"user_id" gorm:"column:user_id"`                                   //该token是颁发给谁的
	UserName              string `json:"username" gorm:"column:username"`                                 //人的名字
	AccessToken           string `json:"access_token" gorm:"column:access_token"`                         //颁发给用户的访问令牌（需要用户携带token来访问接口）
	AccessTokenExpiredAt  int    `json:"access_token_expired_at" gorm:"column:access_token_expired_at"`   //过期时间
	RefreshToken          string `json:"refresh_token" gorm:"column:refresh_token"`                       //刷新令牌
	RefreshTokenExpiredAt int    `json:"refresh_token_expired_at" gorm:"column:refresh_token_expired_at"` //刷新令牌过期时间
	CreatedAt             int64  `json:"created_at" gorm:"column:created_at"`                             //创建时间
	UpdatedAt             int64  `json:"updated_at" gorm:"column:updated_at"`                             //刷新时间
	// AccessTokens          string `json:"access_tokens"`                   //颁发给用户的访问令牌（需要用户携带token来访问接口）
	Role                            user.Roles `gorm:"-"` // 额外补充信息, gorm忽略处理,这样就不需要存在数据库中了
	Access_token_expired_at_seconds string     `gorm:"-"` // 额外补充信息, gorm忽略处理,这样就不需要存在数据库中了
}

func NewToken() *Token {
	return &Token{
		AccessToken:           xid.New().String(),
		AccessTokenExpiredAt:  7200, //多少秒过期 how many seconds expire
		RefreshToken:          xid.New().String(),
		RefreshTokenExpiredAt: 3600 * 24 * 7,
		CreatedAt:             time.Now().Unix(),
		UpdatedAt:             time.Now().Unix(),
		// AccessTokens:          xid.New().String(),

	}
}

// 可以指定该结构体在数据库中对应的表名
func (u *Token) TableName() string {
	return "tokens"
}

// 如果你想将 Token 结构体用于调试或记录日志，可以直接调用 String() 方法，得到一个包含了结构体内容的 JSON 字符串。
func (u *Token) String() string {
	dj, _ := json.Marshal(u)
	return string(dj)
}

// 判断是否过期
func (t *Token) IsExpired() (string, error) {
	// fmt.Println("0-------", time.Now())
	// fmt.Println("0-------", t.ExpiredTime()) //2024-01-24 19:49:35 +0800 CST
	// fmt.Println("0-------", time.Now().Sub(t.ExpiredTime())) //计算当前与传进来的时间的差  -1h44m40.970318s
	duration := time.Since(t.ExpiredTime()) //计算和token 过期的相差时间
	expiredSeconds := duration.Seconds()
	if expiredSeconds > 0 {
		return "", exception.NewNotFound("token: %s 过期了 %f秒", t.AccessToken, duration.Seconds())
	} else {
		return fmt.Sprintf("token: %s 还有 %f 小时过期", t.AccessToken, float64(math.Abs(duration.Hours()))), nil
		//float64(math.Abs(duration.Hours())) 负值转换
	}
}

// 计算token的过期时间
func (t *Token) ExpiredTime() time.Time {
	return time.Unix(t.UpdatedAt, 0).Add(time.Duration(t.AccessTokenExpiredAt) * time.Second)
	// return time.Unix(t.CreatedAt, 0).Add(time.Duration(t.AccessTokenExpiredAt) * time.Second)
	//time.Unix(t.CreatedAt, 0)它接受两个参数，第一个参数是sec int64，第二个参数是nsec int64。
	//time.Duration(t.AccessTokenExpiredAt) * time.Second 将整数值转换为 Go 的时间s。,这里time.Duration(t.AccessTokenExpiredAt)可以直接写数字
}
