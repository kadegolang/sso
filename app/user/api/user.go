package api

import (
	"log"
	"strconv"

	"github.com/kadegolang/sso/app/user"
	"github.com/kadegolang/sso/tools"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (u *UserApiHeadler) Registry(router gin.IRouter) {
	// 在这里注册 Gin 路由
	router.POST("/create/user", u.CreateUserApi)
	router.DELETE("/delete/user", u.DeleteUserApi)
	router.GET("/query/user", u.QueryUserApi)
	router.PATCH("/edit/user", u.PatchUserApi) //put和patch的区别就是，用于put当不存在时会创建

	router.GET("/endpoint1", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from Gin API"})
	})
}

// create user
func (u *UserApiHeadler) CreateUserApi(c *gin.Context) {
	req := user.NewCreateUserRequest()
	c.BindJSON(&req)
	// fmt.Println("ioc-list:", ioc.Config().List())
	// fmt.Println("ioc-count:", ioc.Api().List())
	// a := ioc.Default().Get(impl.AppName).(*impl.UserServiceImpl)
	u.user.CreateUser(c, req)
	c.JSON(200, gin.H{"message": "create user is successful"})
}

// delete user
func (u *UserApiHeadler) DeleteUserApi(c *gin.Context) {
	req := user.NewDeleteUserRequest()
	c.BindJSON(&req)
	err := u.user.DeleteUser(c, req)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{"message": "delete user is successful"})
}

// query user
func (u *UserApiHeadler) QueryUserApi(c *gin.Context) {
	req := user.NewQueryUserRequest()
	if c.Query("username") != "" {
		req.QueryByName = c.Query("username")
	}

	if c.Query("id") != "" {
		id, err := strconv.Atoi(c.Query("id"))
		if err != nil {
			// 处理转换错误
			c.String(400, "Invalid parameter,do no exceed 2147483647 for int32/int")
			return
		}
		req.QueryByID = int64(id)
	}

	ins, err := u.user.QueryUser(c, req)
	switch err {
	case gorm.ErrRecordNotFound:
		c.JSON(400, gin.H{"error": "record not found"})
	default:
		c.JSON(200, gin.H{"message": ins})
	}
}

// edit user
func (u *UserApiHeadler) PatchUserApi(c *gin.Context) {
	req := user.NewEditUserRequest()
	c.BindJSON(&req)
	if req.PassWord == "" {
		c.JSON(400, gin.H{"error": "password is not empty"})
		return
	} else {
		req.PassWord = tools.HashPassword(req.PassWord)
		ins, err := u.user.EditUser(c, req)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(200, gin.H{"message": ins})
		return
	}
}
