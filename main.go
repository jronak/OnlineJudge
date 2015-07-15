package main

import (
	_ "OnlineJudge/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SessionOn = true
	beego.SessionName = "OnlineJudge"
	beego.SessionProvider = "OnlineJudge"
	beego.SessionCookieLifeTime = 0
	beego.SessionProvider = "file"
	beego.SessionSavePath = "./tmp"
	beego.Run()
}
