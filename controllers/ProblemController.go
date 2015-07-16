package controllers

import (
	"OnlineJudge/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	categories = ""
	UpdateCategories()
}

var categories string

type ProblemTypes struct {
	Count      int64
	Categories *orm.ParamsList
}

type ProblemController struct {
	beego.Controller
}

//Caching the categories
func UpdateCategories() {
	p := models.Problem{}
	list, num := p.GetTypes()
	pt := ProblemTypes{Count: num, Categories: list}
	bytes, _ := json.Marshal(&pt)
	categories = string(bytes)
}

/*
  router : /problem/
  Response:
  {"Count":3,"Categories":["Data Structure","","Linked List"]}
*/
func (this *ProblemController) Get() {
	this.Data["json"] = categories
	this.ServeJson()
}

// Json served
// To do : serve in pages, per page 10 problems
func (this *ProblemController) ProblemsByCategory() {
	problemType := this.Ctx.Input.Param(":type")
	problem := models.Problem{Type: problemType}
	problems, _ := problem.GetByType()
	bytes, _ := json.Marshal(problems)
	this.Data["json"] = string(bytes)
	this.ServeJson()
}

// Json served
// To-do : send the name of author as well
func (this *ProblemController) ProblemByStatement() {
	problemName := this.Ctx.Input.Param(":name")
	p := models.Problem{Statement: problemName}
	p.GetByStatement()
	bytes, _ := json.Marshal(&p)
	this.Data["json"] = string(bytes)
	this.ServeJson()
}
