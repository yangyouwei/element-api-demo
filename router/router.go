package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yangyouwei/element-api-demo/api"
	"github.com/yangyouwei/element-api-demo/utils"
)

type RespMessages struct {
	Data interface{} `json:"data"`
	Meta MetaStr	`json:"meta"`
}

type  MetaStr struct {
	Msg string `json:"msg"`
	Status int	`json:"status"`
}


func InitRouter() *gin.Engine {
	router := gin.Default()
	//IndexApi为一个Handler

	v1:=router.Group("/api/v1")
	// 跨域中间件
	v1.Use(utils.CrossDomain())
	//token 校验中间件
	v1.Use(utils.Validate())
	{
		//用户管理

		//登录
		v1.POST("/login", api.POST_Login)
		//获取用户列表
		v1.GET("/users", api.GET_UserList)
		//添加用户
		v1.POST("/users", api.POST_AddUser)
		//根据用户id查询用户信息
		v1.GET("/users/:id", api.GET_SelectUserByID)
		//根据id编辑用户信息
		v1.PUT("/users/:id", api.PUT_UserInfo)
		//删除用户
		v1.DELETE("/users/:id", api.DELETE_UserByID)
		//设置用角色
		v1.PUT("/users/:id/role", api.PUT_SetUserRole)
		//修改用户状态
		v1.PUT("/users/:id/state/:type", api.PUT_UserState)

		//权限管理

	}
	return router
}
