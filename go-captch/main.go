package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserReq struct {
	CodeId string `json:"code_id"`
	Code   string `json:"code"`
}

func Captcha(c *gin.Context) {
	captcha := GenerateCaptcha()
	c.JSON(http.StatusOK, gin.H{
		"data": captcha,
	})
}

func Verify(c *gin.Context) {
	var req UserReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "传入数据有问题",
		})
		return
	}
	captcha := VerifyCaptcha(req.CodeId, req.Code)
	if !captcha {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "验证码有问题",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "验证码正确",
	})
}

func main() {
	r := gin.Default()
	r.GET("/captcha", Captcha)
	r.POST("/verify", Verify)
	err := r.Run(":8888")
	if err != nil {
		panic("启动失败")
	}
}
