package main

import (
	_ "OnlineJudge/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SessionOn = true
	beego.SessionProvider = "file"
	beego.SessionSavePath = "./tmp"
	beego.Run()
}
