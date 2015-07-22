package main

import (
	_ "OnlineJudge/routers"
	"github.com/astaxie/beego"
	"strings"
)

func main() {
	beego.SessionOn = true
	beego.SessionName = "OnlineJudge"
	beego.SessionProvider = "OnlineJudge"
	beego.SessionCookieLifeTime = 0
	beego.SessionProvider = "file"
	beego.SessionSavePath = "./tmp"

	beego.AddFuncMap("n2br", n2br)
	beego.Run()
}

func n2br(str string) string {
	return strings.Replace(str,"\n","<br/>",-1)
}
