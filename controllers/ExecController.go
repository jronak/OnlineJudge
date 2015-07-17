package controllers

import (
	"OnlineJudge/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type ExecController struct {
	beego.Controller
}

/* Expects Json of format shown below - Ajax request
`{"Program":{"Lang":"C","Stdin":"If any custom input by user"},"RawCode":"Code should be here",
	"Pid":12,"IsCustomInput":true,"IsBatch":false}`
 Response:
 	{"Lang":"C","Stderr":"Error if any","Stdout":"Output of the program",
		"CompilationStatus":false,"RunStatus":false,"ExecTime":0}
*/

// Input is manully taken from the form as now.
//Later Input will be accpted as expected. 6 is pid for the sake of testing
func (this *ExecController) Post() {
	c := models.Exec(6, this.GetString("code"), this.GetString("language"), this.GetString("stdin"))
	data, _ := json.Marshal(&c)
	this.Data["json"] = string(data)
	this.ServeJson()
}

//Get method to Exec will casue redirection to home page
func (this *ExecController) Get() {
	this.TplNames = "testExec.html"
}

//Execpts only Post requests!
func (this *ExecController) Submit() {
	c := models.SubmitUpdateScore(2, 10, this.GetString("code"), this.GetString("language"))
	data, _ := json.Marshal(&c)
	this.Data["json"] = string(data)
	this.ServeJson()
}
