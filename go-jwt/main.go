package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type UserReq struct {
	Username string `json:"username"`
	Password string `json:""`
}

func Login(c *gin.Context) {
	var userinfo UserReq
	err := c.ShouldBind(&userinfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "传入数据有问题",
		})
		return
	}
	if userinfo.Username != "admin" || userinfo.Password != "123456" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "账号或者密码有问题",
		})
		return
	}
	token, err := GenerateToken(userinfo.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "token生成失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": token,
	})
}

func Userinfo(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	if len(auth) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户未登录",
		})
		return
	}
	token := strings.Split(auth, " ")[1]
	parseToken, err := ParseToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "token解析失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": parseToken.Username,
	})
}

func main() {
	r := gin.Default()
	r.POST("/login", Login)
	r.GET("/userinfo", Userinfo)
	err := r.Run(":8080")
	if err != nil {
		panic("请求失败")
	}
}
