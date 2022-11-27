package api

import (
	"gin-demo/api/middleware"
	"gin-demo/dao"
	"gin-demo/model"
	"gin-demo/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func getUsernameFromToken(c *gin.Context) {
	username, _ := c.Get("username")
	utils.RespSuccess(c, username.(string))
}

func register(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespSuccess(c, "请输入账号或密码")
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")
	question := c.PostForm("question")
	answer := c.PostForm("answer")

	flag := dao.SelectUser(username)

	if flag {
		utils.RespFailed(c, "此账号已存在")
		return
	}
	dao.AddUser(username, password)
	dao.AddAnswer(question, answer)
	utils.RespSuccess(c, "注册账号成功")
}

func login(c *gin.Context) {
	if err := c.ShouldBind(&model.User{}); err != nil {
		utils.RespFailed(c, "请输入账号和密码")
		return
	}

	username := c.PostForm("username")
	password := c.PostForm("password")

	flag := dao.SelectUser(username)

	if !flag {
		utils.RespFailed(c, "此账号不存在")
		return
	}

	selectPassword := dao.SelectPasswordFromUsername(username)

	if selectPassword != password {
		utils.RespFailed(c, "密码错误")
		return
	}

	claim := model.MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    "ZYQ",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claim)
	tokenString, _ := token.SignedString(middleware.Secret)
	utils.RespSuccess(c, "登陆成功"+tokenString)
}
