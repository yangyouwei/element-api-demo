package utils

import (
	"crypto/sha256"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yangyouwei/element-api-demo/api"
	"golang.org/x/crypto/pbkdf2"
	logger "log"
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
			c.AbortWithStatus(api.StatusNoContent)
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
			c.JSON(api.StatusUnauthorized,gin.H{"message":"身份验证失败"})
		}
	}
}

//jwt token验证
func AuthToken(token string) bool {
	atoken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjUwMCwicmlkIjowLCJpYXQiOjE1MTI1NDQyOTksImV4cCI6MTUxMjYzMDY5OX0.eGrsrvwHm-tPsO9r_pxHIQ5i5L1kX9RX444uwnRGaIM"
	if token == atoken {
		return true
	}else {
		return false
	}
}

//生成jwt token
func GenToken() string {
	return "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOjUwMCwicmlkIjowLCJpYXQiOjE1MTI1NDQyOTksImV4cCI6MTUxMjYzMDY5OX0.eGrsrvwHm-tPsO9r_pxHIQ5i5L1kX9RX444uwnRGaIM"
}

//密码加密
func PasswordEncrypt(p string) string {
	//生成随机盐
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		salt = []byte{219,152,152,26,33,17,4,26,126,174,154,135,175,53,193,26,179,14,163,28,138,88,187,138,115,224,98,253,12,215,110,1}
		logger.Println(err)
	}
	//生成密文
	dk := pbkdf2.Key([]byte(p), salt, 2, 32, sha256.New)
	return fmt.Sprint(hex.EncodeToString(dk))
}
