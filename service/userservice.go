package service

import (
	"strconv"

	"github.com/Withmm/IM/models"
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
	user.PassWord = password

	if err := validator.New().Struct(user); err != nil {
		c.JSON(400, gin.H{
			"input error": "User attribute input error:" + err.Error(),
		})
		return
	}

	user.PassWord = password
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
