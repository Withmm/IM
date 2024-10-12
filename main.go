package main

import (
	"github.com/Withmm/IM/router"
	"github.com/Withmm/IM/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()
	r := router.Router()
	r.Run()
}
