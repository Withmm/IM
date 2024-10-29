package service

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Withmm/IM/models"
	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles("index.html", "views/chat/head.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "index")
}

func ToChat(c *gin.Context) {
	//fmt.Println("call Tochat.............................")
	ind, err := template.ParseFiles("views/chat/index.html",
		"views/chat/head.html",
		"views/chat/foot.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/createcom.html",
		"views/chat/userinfo.html",
		"views/chat/main.html")
	if err != nil {
		panic(err)
	}
	user := models.UserBasic{}
	userId, _ := strconv.Atoi(c.Query("userId"))
	token := c.Query("token")
	user.ID = uint(userId)
	user.Identity = token
	ind.Execute(c.Writer, user)
}

func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("./views/user/register.html")
	if err != nil {
		c.String(http.StatusInternalServerError, "页面加载失败，请稍后重试")
		return
	}
	ind.Execute(c.Writer, "register")
}

func ToLogin(c *gin.Context) {
	ind, err := template.ParseFiles("login.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, nil)
}
