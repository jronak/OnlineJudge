package routers

import (
	"OnlineJudge/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// List problems (or have a seperate controller as IndexController?)
	beego.Router("/", &controllers.ProblemController{}, "*:List")

	// On User -- Includes user profiles, settings and login and sign up routes
	beego.AutoRouter(&controllers.UserController{})

	// On Problems
	beego.Router("/problem/create", &controllers.ProblemController{}, "*:Create;post:SaveProblem")
	beego.Router("/problem/:type", &controllers.ProblemController{}, "*:ProblemsByCategory")
	beego.Router("/problem/:type/:id", &controllers.ProblemController{}, "*:ProblemByStatement")
	beego.Router("/problem/:type/:id/submit", &controllers.ProblemController{}, "post:SaveSubmission;*:Submit") // ->ProblemController(notes that user has tried solving problem)->ExecController(seek for helper to exec)->ProblemController(get result info & build on it)
	beego.Router("/problem/:type/:id/edit", &controllers.ProblemController{}, "post:SaveProblem;*:Edit")
}
