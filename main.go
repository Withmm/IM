package main

import (
	"github.com/Withmm/IM/router"
	"github.com/Withmm/IM/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()
	r := router.Router()
	r.Run()
}
