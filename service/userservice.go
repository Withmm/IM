package service

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Withmm/IM/models"
	"github.com/Withmm/IM/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
)

func GetUserList(c *gin.Context) {
	data := models.GetUserList()

	c.JSON(200, gin.H{"message": data})
}

func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.PostForm("name")
	//user.Email = c.PostForm("email")
	//user.Phone = c.PostForm("phone")
	password := c.PostForm("password")
	repassword := c.PostForm("Identity")

	if password != repassword {
		c.JSON(200, gin.H{
			"input error": "The two passwords do not match",
		})
		return
	}
	// 将明文储存为密文
	salt, err := utils.GenerateSalt(16)
	if err != nil {
		c.JSON(200, gin.H{
			"salt generate error": err.Error(),
		})
	}
	user.Salt = salt
	user.PassWord = utils.MakePassword(password, user.Salt)

	//输入信息校验
	if err := validator.New().Struct(user); err != nil {
		c.JSON(200, gin.H{
			"input error": "User attribute input error:" + err.Error(),
		})
		return
	}

	//调用后端修改数据库
	if err := models.CreateUser(user); err != nil {
		c.JSON(200, gin.H{
			"database error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully"})
}

func DeleteUser(c *gin.Context) {
	// 获取被删除对象的id
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}

	if err := models.DeleteUser(idInt); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	c.JSON(200, gin.H{
		"message": "delete user (id = " + id + ") successfully",
	})
}

func UpdateUser(c *gin.Context) {
	// 获取更新对象的id
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	var req map[string]interface{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	if err := models.UpdateUser(idInt, req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "user update successfully",
	})
}

func FindUserByNameAndPassword(c *gin.Context) {
	// 在数据库中查找
	name := c.PostForm("name")
	password := c.PostForm("password")
	user, err := models.FindUserByNameAndPassword(name, password)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": err.Error(),
			"user":    user,
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    0,
		"message": "Login successful",
		"user":    user,
	})
	/*
		c.JSON(200, gin.H{
			"message": "user found successfully",
		})
	*/
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	// fmt.Println("协议准备升级为websocket")
	// 首先将协议升级为websocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("websocket升级完成")
	// 生命周期管理
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(ws)

	MsgHandle(ws, c)
}

func MsgHandle(ws *websocket.Conn, c *gin.Context) {
	// fmt.Println("MsgHandle...")
	for {
		msg, err := utils.Subscibe(c, utils.PublishKey)
		if err != nil {
			fmt.Println(err)
			return
		}
		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]: %s", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}
