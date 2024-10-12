package service

import (
	"strconv"

	"github.com/Withmm/IM/models"
	"github.com/Withmm/IM/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetUserList(c *gin.Context) {
	data := models.GetUserList()

	c.JSON(200, gin.H{"message": data})
}

func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.PostForm("name")
	user.Email = c.PostForm("email")
	user.Phone = c.PostForm("phone")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")

	if password != repassword {
		c.JSON(400, gin.H{
			"input error": "The two passwords do not match",
		})
		return
	}
	// 将明文储存为密文
	salt, err := utils.GenerateSalt(16)
	if err != nil {
		c.JSON(400, gin.H{
			"salt generate error": err.Error(),
		})
	}
	user.Salt = salt
	user.PassWord = utils.MakePassword(password, user.Salt)

	//输入信息校验
	if err := validator.New().Struct(user); err != nil {
		c.JSON(400, gin.H{
			"input error": "User attribute input error:" + err.Error(),
		})
		return
	}

	//调用后端修改数据库
	if err := models.CreateUser(user); err != nil {
		c.JSON(400, gin.H{
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
	name := c.PostForm("name")
	password := c.PostForm("password")
	if err := models.FindUserByNameAndPassword(name, password); err != nil {
		c.JSON(400, gin.H{
			"data error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "user found successfully",
	})
}
