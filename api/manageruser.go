package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginUser struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

//{
//"data": {
//"id": 500,
//"rid": 0,
//"username": "admin",
//"mobile": "123",
//"email": "123@qq.com",
//"token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpxHIQ5i5L1kX9RX444uwnRGaIM"
//},
//"meta": {
//"msg": "登录成功",
//"status": 200
//}
//}

//login
func POST_Login(c *gin.Context)  {

}


//{
//"data": {
//"totalpage": 5,
//"pagenum": 4,
//"users": [
//{
//"id": 25,
//"username": "tige117",
//"mobile": "18616358651",
//"type": 1,
//"email": "tige112@163.com",
//"create_time": "2017-11-09T20:36:26.000Z",
//"mg_state": true, // 当前用户的状态
//"role_name": "炒鸡管理员"
//}
//]
//},
//"meta": {
//"msg": "获取成功",
//"status": 200
//}
//}

//获取用户列表,分页显示
func GET_UserList(c *gin.Context)  {

}


//{
//"data": {
//"id": 28,
//"username": "tige1200",
//"mobile": "test",
//"type": 1,
//"openid": "",
//"email": "test@test.com",
//"create_time": "2017-11-10T03:47:13.533Z",
//"modify_time": null,
//"is_delete": false,
//"is_active": false
//},
//"meta": {
//"msg": "用户创建成功",
//"status": 201
//}
//}
//添加用户
func POST_AddUser(c *gin.Context)  {

}

//设置用户状态
func PUT_UserState(c *gin.Context)  {
	id := c.Param("id")
	fmt.Println("id: ",id)
	state := c.Param("type")
	fmt.Println("type: ",state)
}

//根据id查询用户信息
func GET_SelectUserByID(c *gin.Context)  {
	id := c.Param("id")
	fmt.Println(id)
}

//编辑用户信息
func PUT_UserInfo(c *gin.Context)  {

}

//根据id删除用户
func DELETE_UserByID(c *gin.Context)  {

}

//设置用户角色
func PUT_SetUserRole(c *gin.Context)  {

}
