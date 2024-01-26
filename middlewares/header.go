package middlewares

import "github.com/gin-gonic/gin"

// 解决cors跨域请求
func CorsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, Authorization")

	// 如果是预检请求（OPTIONS），直接返回 200
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	// 继续处理请求
	c.Next()
}
