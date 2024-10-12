package main

import (
	"github.com/Withmm/IM/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:200510@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("fail to connect database")
	}

	db.AutoMigrate(&models.UserBasic{})

	// create
	//user := &models.UserBasic{}
	//user.Name = "xiongzile"
	//db.Create(user)

	// read
	//fmt.Println(db.First(user, 1))

	// update
	//db.Model(user).Update("PassWord", "1234")

}
