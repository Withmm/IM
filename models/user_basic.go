package models

import (
	"errors"
	"fmt"

	"github.com/Withmm/IM/utils"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string `validate:"required,min=2,max=100"`
	PassWord      string `validate:"required,min=6,max=24"`
	Phone         string `validate:"required,len=11"`
	Email         string `validate:"required,email"`
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     uint64
	HeartBeattime uint64
	LogOutTime    uint64
	IsLogOut      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)

	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func CreateUser(user UserBasic) error {
	tmp := UserBasic{}

	if result := utils.DB.Where("name = ?", user.Name).First(&tmp); result.Error == nil {
		return fmt.Errorf("user : %s has been registered", user.Name)
	}

	if result := utils.DB.Where("email = ?", user.Email).First(&tmp); result.Error == nil {
		return fmt.Errorf("email : %s has been registered", user.Email)
	}

	if result := utils.DB.Where("phone = ?", user.Phone).First(&tmp); result.Error == nil {
		return fmt.Errorf("phone number : %s****%s has been registered", user.Phone[0:3], user.Phone[7:11])
	}

	return utils.DB.Create(&user).Error
}

func DeleteUser(id int) error {
	// 根据id找到用户
	user := UserBasic{}
	result := utils.DB.First(&user, id)

	// check if id exists
	if result.Error != nil {
		return errors.New("user not found")
	}

	return utils.DB.Delete(&user).Error
}

func UpdateUser(id int, updates map[string]interface{}) error {
	// 根据id查找到需要更新的用户
	user := UserBasic{}
	result := utils.DB.First(&user, id)

	if result.Error != nil {
		return errors.New("user not found")
	}

	return utils.DB.Model(&user).Updates(updates).Error
}
