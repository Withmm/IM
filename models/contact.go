package models

import (
	"fmt"

	"github.com/Withmm/IM/utils"
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	OwnerId  uint   // whose contact information
	TargetId uint   // to who
	Type     int    // 1, 2, 3  1->friend 2->group 3->reserve
	Desc     string // TODO
}

func (table *Contact) TableName() string {
	return "contact"
}

func GetFriends(ownerID uint) []UserBasic {
	//TODO
	return []UserBasic{}
}

func GetGroups(ownerID uint) []GroupBasic {
	//TODO
	return []GroupBasic{}
}

func AddFriend(ownerID uint, targetID uint) error {
	//首先检查添加的对象在不在好友列表里
	user := UserBasic{}
	result := utils.DB.First(&user, targetID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("用户 %d 不存在", targetID)
		}
	}
	if user.ID == ownerID {
		return fmt.Errorf("请不要添加自己为好友")
	}
	ctact := &Contact{
		OwnerId:  ownerID,
		TargetId: targetID,
	}

	result = utils.DB.Create(ctact)
	if result.Error != nil {
		return result.Error // 返回插入错误
	}

	return nil
}
