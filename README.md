## 之前的ioc
map形式，注册name：“interface" 如果掉用其方法，就是实现init，这样才调

## 现在的ioc
切片，只需要注册 name，version
如果掉用其方法，就是实现init，这样才调，但是这种方式只要name和version，比上面更复杂好用


### app1 ioc1 是老的版本  map
### app  ioc  是新的版本  slice



## 路由注册

## func main() {

	req := ioc.NewLoadConfigRequest()
	req.ConfigFile.Enabled = true
	req.ConfigFile.Path = "etc/application.toml"
	ioc.DevelopmentSetup(req)
	// a := ioc.Default().Get(impl.AppName).(*impl.UserServiceImpl)
	// req1 := user.NewCreateUserRequest()
	// req1.UserName = "sja11"
	// req1.PassWord = "sman11"
	// req1.Label = map[string]string{"bsa1b": "111111"}
	// t.Log(a.CreateUser(context.Background(), req1))
	// 加载 Gin API，指定路径前缀和路由器对象
	r := gin.Default()
	// // a := r.Group("ccc")
	// r.POST("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "Hello from Gin API"})
	// })
	r.LoadHTMLGlob("views/*") //配置模版中间件
	ioc.LoadGinApi("/api/v1", r)
	// 启动 Gin 服务器
	r.Run(":8080")
	}


## api路由注册需要满足以下条件

### api的实现方法要先注册api ioc，在注册路由api

	func init() {
		ioc.Api().Registry(&TokenApiHeadler{})
	}

	type TokenApiHeadler struct {
		//因为这个api 依赖他的具体实现
		token token.Service
		// user *impl.UserServiceImpl //依赖控制器 不应该这样干他的实现，应该user.Service方法才对

	}

	func (u *TokenApiHeadler) Init() error {
		u.token = ioc.Default().Get(tokenimpl.AppName).(*tokenimpl.TokenServiceImpl)
		return nil
	}

	func (u *TokenApiHeadler) Name() string {
		return "tokens"
	}

	func (u *TokenApiHeadler) Version() string {
		return ""
	}

	func (u *TokenApiHeadler) Priority() int {
		return 9
	}

	func (u *TokenApiHeadler) Close(ctx context.Context) error {
		return nil
	}

	func (u *TokenApiHeadler) AllowOverwrite() bool {
		return true
	}


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
	})}
