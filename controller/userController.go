package controller

import (
	"douyin/pojo"
	"douyin/service"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	pojo.Response
}

var USI service.userServiceImpl

// localhost:8080/register post方法?username=&password="
func Register(c *gin.Context) {
	// 从路径变量中获取参数
	userName := c.Query("username")
	password := c.Query("password")
	c.JSON(200, gin.H{"username": userName, "password": password})

	//user := USI.GetUserByUserName(userName)
	//if user.UserName == userName {
	//	fmt.Println("用户名已存在")
	//}
	////创建新一个用户
	//USI.CreatUser(dao.UserDao{
	//	UserName: userName,
	//	Password: password,
	//})

}
