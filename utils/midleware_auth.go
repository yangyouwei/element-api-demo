package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var TokensMap = make(map[string]string)

//跨域中间件
func CrossDomain() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

//token校验中间件
func Validate() gin.HandlerFunc{
	return func(c *gin.Context){
		token := c.Request.Header.Get("Authorization")
		if c.Request.URL.Path == "/api/v1/login" {
			c.Next()
		}else if AuthToken(token) {
			c.Next()
		} else {
			c.Abort()
			c.JSON(http.StatusUnauthorized,gin.H{"message":"身份验证失败"})
		}
	}
}

//if token == "" {
//c.JSON(http.StatusOK,gin.H{"message":"身份验证成功"})
//c.Next()  //该句可以省略，写出来只是表明可以进行验证下一步中间件，不写，也是内置会继续访问下一个中间件的
//}else {
//c.Abort()
//c.JSON(http.StatusUnauthorized,gin.H{"message":"身份验证失败"})
//return// return也是可以省略的，执行了abort操作，会内置在中间件defer前，return，写出来也只是解答为什么Abort()之后，还能执行返回JSON数据
//}

//jwt token验证
func AuthToken(token string) bool {
	if false {
		return true
	}else {
		return false
	}
}

//生成token
func GenToken() string {
	return ""
}