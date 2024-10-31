package router

import (
	"github.com/Withmm/IM/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	e := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	e.Use(sessions.Sessions("session", store))
	// homepage
	e.GET("/", service.GetIndex)
	e.GET("/index", service.GetIndex)
	//登录和注册界面
	e.GET("/user/toLogin", service.ToLogin)
	e.GET("/user/toRegister", service.ToRegister)

	//登录（账号密码检查）
	e.POST("/user/login", service.Login)
	// 注册
	e.POST("/user/register", service.Register)

	//主界面的三大子页面
	e.GET("/toChatPage", service.ToChatPage)
	e.GET("/toFriendPage", service.ToFriendPage)
	e.GET("/toProfilePage", service.ToProfilePage)

	e.POST("/addFriend", service.AddFriend)
	e.POST("/removeFriend", service.RemoveFriend)
	// delete user by id
	e.DELETE("/user/:id", service.DeleteUser)
	// update user by id
	e.PUT("/user/:id", service.UpdateUser)

	//msg module
	e.GET("/user/sendMsg", service.SendMsg)
	e.GET("/user/sendUserMsg", service.SendUserMsg)

	e.GET("/toChat", service.ToChat)
	return e
}
