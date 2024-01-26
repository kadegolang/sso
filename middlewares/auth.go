package middlewares

import (
	"fmt"
	"net/http"

	"gitee.com/go-kade/library/exception"
	"gitee.com/go-kade/library/ioc"
	"github.com/kadegolang/sso/app/token"
	tokenimpl "github.com/kadegolang/sso/app/token/impl"
	"github.com/gin-gonic/gin"
)

// 用于鉴权的中间件
// 用于Token鉴权的中间件
type TokenAuth struct {
	tk token.Service
	// role user.Role
}

func NewTokenAuther() *TokenAuth {
	return &TokenAuth{
		// tk: ioc.Default().Get(tokenimpl.AppName).(*tokenimpl.TokenServiceImpl),
		tk: ioc.Default().Get(tokenimpl.AppName).(token.Service),
	}
}

// 怎么鉴权?
// Gin中间件 func(*Context)
func (a *TokenAuth) Auth(g *gin.Context) {
	//1.获取token
	//g.GetHeader("") 获取cookis中的access_token两种方式,登录时cookis携带了用户access_token，这里是拿到那个token做鉴权
	web_cookis, err := g.Cookie(token.TOKEN_COOKIE_NAME)
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Printf("error code:%d\nerror message:%s\n", exception.NewWebCookisNotFoundnRequest("").ErrCode, exception.NewWebCookisNotFoundnRequest("").Reason)
			g.Redirect(302, "/api/v1/tokens/login1")
			g.Abort()
			// g.JSON(exception.NewWebCookisNotFoundnRequest("").HttpCode, exception.NewWebCookisNotFoundnRequest("").Reason)
			return
		}
		// return
	}
	fmt.Println("------", web_cookis)
	//2.验证token是否过期
	req := token.NewValiateTokenRequst(web_cookis)
	tk, err := a.tk.ValiateToken(g.Request.Context(), req)
	if err != nil {
		fmt.Printf("error code:%d\nerror message:%s\n", exception.NewWebCookisNotFoundnRequest("").ErrCode, exception.NewWebCookisNotFoundnRequest("").Reason)
		g.Redirect(302, "/api/v1/tokens/login1")
		g.Abort()
		return
	}
	//3.把token放到gin context 上下文中
	fmt.Println("---------tk:", tk)
	if g.Keys == nil {
		g.Keys = map[string]any{}
	}
	g.Keys[token.TOKEN_GIN_KEY_NAME] = tk
	// 后面从上下文中调取
	// tkobj := g.Keys[token.TOKEN_GIN_KEY_NAME]
	// tkobj1 := tkobj.(*token.Token)
	// fmt.Println("------sanskanskana", tkobj1.AccessToken)
}
