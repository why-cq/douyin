package controller

import (
	"douyin/dao"
	"douyin/pojo"
	"douyin/service"
	"douyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserResponse struct {
	pojo.Response
	UserId     int64  `json:"user_id"`
	TokenValue string `json:"token"`
}
type UserInf struct {
	pojo.Response
	UserId        int64  `json:"id"`
	UserName      string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

var USI service.UserServiceImpl
var FSI service.FollowServiceImpl

// localhost:8080/register post方法?username=&password="
func Register(c *gin.Context) {
	// 从路径变量中获取参数
	var token utils.Token
	userName := c.Query("username")
	password := c.Query("password")
	//c.JSON(200, gin.H{"username": userName, "password": password})
	//1.用户名已存在
	user := USI.GetUserByUserName(userName)
	if user.UserName == userName {
		fmt.Println("用户名已存在")
		c.JSON(200, UserResponse{
			Response: pojo.Response{1, "用户名已存在"},
		})
	} else {
		// 用户不存在
		//创建新一个用户
		ok := USI.CreatUser(dao.UserDao{
			UserName: userName,
			Password: password,
		})
		if ok {
			// 返回的第一个值通过操作做数据库获取
			user1, _ := dao.GetUserByUsername(userName)
			// token通过工具类生成
			token, _ = utils.GenerateToken(userName, password)

			c.JSON(200, UserResponse{
				Response:   pojo.Response{0, "注册用户成功"},
				UserId:     user1.Id,
				TokenValue: token.Values,
			})
		} else {
			c.JSON(200, gin.H{"message": "创建用户失败"})
		}
	}
}

// localhost:8080/Login post方法?username=&password="
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	var token utils.Token

	user := USI.GetUserByUserName(username)
	//1.用户名不存在
	if user.UserName != username {
		fmt.Println("用户名不存在")
		c.JSON(200, gin.H{"msg": "用户名不存在"})
	} else {
		// 用户存在, 比对密码是否正确
		if user.Password != password {
			fmt.Println("密码错误")
			c.JSON(200, gin.H{"msg": "密码错误"})
		} else {
			//生成token
			token, _ = utils.GenerateToken(username, password)
			c.JSON(200, UserResponse{
				Response:   pojo.Response{0, "登录成功"},
				UserId:     user.Id,
				TokenValue: token.Values,
			})

		}

	}

}

// localhost:8080/user?user_id=2 get方法?user_id=""&token=""
func UserInfo(c *gin.Context) {
	//解析url中的参数转换为int64
	u := c.Query("user_id")
	user_id, _ := strconv.ParseInt(u, 10, 64)
	user := USI.GetUserById(user_id)
	// 关注总数
	followCount, _ := FSI.GetFollowersByUserId(user_id)
	// 粉丝总数
	followerCount, _ := FSI.GetFansByUserId(user_id)

	// 是否关注怎么个业务逻辑?????
	c.JSON(http.StatusOK, UserInf{
		Response:      pojo.Response{0, "获取用户信息成功"},
		UserId:        user.Id,
		UserName:      user.UserName,
		FollowCount:   followCount,
		FollowerCount: followerCount,
		IsFollow:      false,
	})

}
