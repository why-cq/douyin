package controller

import (
	"douyin/dao"
	"douyin/pojo"
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	pojo.Response
}

//var USI service.UserServiceImpl

// localhost:8080/register post方法?username=&password="
func Register(c *gin.Context) {
	// 从路径变量中获取参数
	userName := c.Query("username")
	password := c.Query("password")
	//c.JSON(200, gin.H{"username": userName, "password": password})
	USI := service.UserServiceImpl{}
	//1.用户名已存在
	user := USI.GetUserByUserName(userName)
	if user.UserName == userName {
		fmt.Println("用户名已存在")
		c.JSON(200, gin.H{"message": "用户已存在"})
	} else {
		// 用户不存在
		//创建新一个用户
		ok := USI.CreatUser(dao.UserDao{
			UserName: userName,
			Password: password,
		})
		if ok {
			c.JSON(200, gin.H{"username": userName, "password": password})
		}

		c.JSON(200, gin.H{"message": "创建用户失败"})

	}

}
