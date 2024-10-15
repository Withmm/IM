package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	OwnerId  uint   // whose contact information
	TargetId uint   // to who
	Type     int    // 0, 1, 2
	Desc     string // TODO
}

func (table *Contact) TableName() string {
	return "contact"
}
