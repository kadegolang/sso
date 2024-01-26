package api

import (
	"fmt"
	"net/http"

	"gitee.com/go-kade/library/exception"
	"github.com/kadegolang/sso/app/token"
	"github.com/kadegolang/sso/middlewares"
	"github.com/gin-gonic/gin"
)

func (t *TokenApiHeadler) Registry(router gin.IRouter) {
	///解决cors跨域请求
	router.Use(middlewares.CorsMiddleware)
	router.GET("/login1", func(g *gin.Context) {
		// 使用 HTML 渲染登录页面
		g.HTML(http.StatusOK, "login.html", nil)
	})
	router.POST("/login", t.LoginApi)
	router.OPTIONS("/login", t.LoginApi)
	router.GET("/loginout", t.LoginOutApi)
	// 后台管理接口 需要认证  这个有个规范，必须把需要修改的请求放在下面
	router.Use(middlewares.NewTokenAuther().Auth)
	//假设页面调用
	router.GET("/yw", func(g *gin.Context) {
		// 使用 HTML 渲染登录页面
		g.HTML(http.StatusOK, "login.html", nil)
	})
}

func (t *TokenApiHeadler) LoginApi(g *gin.Context) {
	//0.移除浏览器上面的cookies
	g.SetCookie(token.TOKEN_COOKIE_NAME, "", -1, "/", "localhost", false, true)

	//1.解析请求参数object
	req := token.NewLoginRequst()

	//定义一个请求类型
	content := g.GetHeader("Content-Type")
	switch content {
	case "application/json":
		if err := g.BindJSON(&req); err != nil {
			g.JSON(exception.NewBadRequest(err.Error()).ErrCode, exception.NewBadRequest(err.Error()))
			return
		}
	case "application/x-www-form-urlencoded":
		req.Username = g.PostForm("username")
		req.Password = g.PostForm("password")
	default:
		// 未知的 Content-Type
		g.JSON(400, gin.H{"error": "Unsupported Content-Type"})
		return
	}

	//2.执行登陆逻辑
	// g.Request.Context() http请求的context
	in, err := t.token.Login(g.Request.Context(), req)
	if err != nil {
		g.JSON(exception.NewBadRequest(err.Error()).ErrCode, exception.NewBadRequest(err.Error()))
		return
	}
	// access_token 通过SetCookie 直接写到浏览器客户端(Web)
	//这种方式直接设置cookis
	g.SetCookie(token.TOKEN_COOKIE_NAME, in.AccessToken, 0, "/", "localhost", false, true)
	// 3. 返回响应
	g.JSON(http.StatusOK, in)
}

func (t *TokenApiHeadler) LoginOutApi(g *gin.Context) {
	//1.获取浏览器端端cookis
	web_cookis, err := g.Cookie(token.TOKEN_COOKIE_NAME)
	if err != nil {
		if err == http.ErrNoCookie {
			g.JSON(exception.NewWebCookisNotFoundnRequest("").HttpCode, exception.NewWebCookisNotFoundnRequest("").Reason)
			return
		}
		g.JSON(exception.NewWebCookisNotFoundnRequest("").HttpCode, exception.NewWebCookisNotFoundnRequest("").Reason)
		return
	}

	//2.创建*token.LoginOutRequst 请求体
	req := token.NewLoginOutRequst()
	req.AccessToken = web_cookis

	//3.调用退出umpl
	err = t.token.Logout(g.Request.Context(), req)
	if err != nil {
		g.JSON(400, err)
	}
	//4.移除浏览器上面的cookies
	g.SetCookie(token.TOKEN_COOKIE_NAME, "", -1, "/", "localhost", false, true)
	g.JSON(200, fmt.Sprintf("%s cookis is remove", req.AccessToken))
}
