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

func GetFriends(ownerID uint) ([]UserBasic, error) {
	var friends []UserBasic

	// get contact item
	var contacts []Contact
	if err := utils.DB.Where("owner_id = ? AND type = 1", ownerID).Find(&contacts).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// if no friend, it's not really a fault.
			return nil, nil
		}
		return nil, err
	}
	// get user profile from contact targetID
	for _, contact := range contacts {
		var friend UserBasic
		id := contact.TargetId
		if err := utils.DB.First(&friend, id).Error; err != nil {
			return nil, err
		}
		friends = append(friends, friend)
	}
	return friends, nil
}

func GetGroups(ownerID uint) []GroupBasic {
	//TODO
	return []GroupBasic{}
}

func AddFriend(ownerID uint, targetID uint) error {
	//check if targetId is right.
	user := UserBasic{}
	result := utils.DB.First(&user, targetID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("用户 %d 不存在", targetID)
		}
	}

	// check if user is adding themselves.
	if user.ID == ownerID {
		return fmt.Errorf("请不要添加自己为好友")
	}

	// check if adding twice.
	contactTmp := &Contact{}
	if err := utils.DB.Where("owner_id = ? AND target_id = ?", ownerID, targetID).First(contactTmp).Error; err == nil {
		return fmt.Errorf("您已经添加过该用户了")
	}

	ctact := &Contact{
		OwnerId:  ownerID,
		TargetId: targetID,
		Type:     1, //friend contact
	}

	result = utils.DB.Create(ctact)
	if result.Error != nil {
		return result.Error // 返回插入错误
	}

	return nil
}
func RemoveFriend(ownerID uint, targetID uint) error {
	//check if targetId is right.
	if ownerID == targetID {
		return fmt.Errorf("don't try to remove yourself from a friend list")
	}
	user := UserBasic{}
	result := utils.DB.First(&user, targetID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("user %d does not exists", targetID)
		}
		return result.Error
	}
	contact := &Contact{
		OwnerId:  ownerID,
		TargetId: targetID,
	}
	delResult := utils.DB.Where("owner_id = ? AND target_id = ?", ownerID, targetID).Delete(contact)
	if delResult.Error != nil {
		return delResult.Error
	}
	if delResult.RowsAffected == 0 {
		return fmt.Errorf("friendship between user %d and user %d does not exists", ownerID, targetID)
	}
	return nil
}
