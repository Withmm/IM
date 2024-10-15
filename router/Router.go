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
	e.GET("/user/userList", service.FindUserByNameAndPassword)
	// create new user
	e.POST("/user/userList", service.CreateUser)
	// delete user by id
	e.DELETE("/user/:id", service.DeleteUser)
	// update user by id
	e.PUT("/user/:id", service.UpdateUser)

	//msg module
	e.GET("/user/sendMsg", service.SendMsg)
	e.GET("/user/sendUserMsg", service.SendUserMsg)
	return e
}
