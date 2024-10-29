package router

import (
	"github.com/Withmm/IM/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	e := gin.Default()

	e.Static("/asset", "./asset/")
	e.LoadHTMLGlob("views/**/*")
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
