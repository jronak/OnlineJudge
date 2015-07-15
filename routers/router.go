package routers

import (
	"OnlineJudge/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//Exec is running at the root
	beego.Router("/", &controllers.ExecController{}, "*:Get;post:Post")
	beego.Router("/submit", &controllers.ExecController{}, "post:Submit;*:Get")
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/problem", &controllers.ProblemController{})
	beego.Router("/problem/:type", &controllers.ProblemController{}, "*:ProblemsByCategory")
	beego.Router("/problem/:type/:name", &controllers.ProblemController{}, "*:ProblemByStatement")
}
