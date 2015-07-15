package routers

import (
	"OnlineJudge/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ExecController{}, "*:Get;post:Post")
	beego.Router("/submit", &controllers.ExecController{}, "post:Submit;*:Get")
}
