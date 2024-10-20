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
	// check userlist
	e.POST("/user/findUserByNameAndPwd", service.FindUserByNameAndPassword)
	// create new user
	e.POST("/user/createUser", service.CreateUser)
	// delete user by id
	e.DELETE("/user/:id", service.DeleteUser)
	// update user by id
	e.PUT("/user/:id", service.UpdateUser)

	//msg module
	e.GET("/user/sendMsg", service.SendMsg)
	e.GET("/user/sendUserMsg", service.SendUserMsg)

	e.GET("/toRegister", service.ToRegister)
	e.GET("/toChat", service.ToChat)
	return e
}
