package controllers

import (
	"OnlineJudge/models"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"strconv"
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
	BaseController
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

// Create Page
func (this *ProblemController) Create() {
	this.Data["title"] = "Create Problem "

	this.Layout = "layout.tpl"
	this.TplNames = "problem/create.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = ""
	this.LayoutSections["Sidebar"] = ""
}

// Save Problem
// To-do: Clean info before save
// To-do: Check login and user previlages
func (this *ProblemController) SaveProblem() {
	points, _ := strconv.Atoi(this.GetString("points"))
	problem := models.Problem{
		Statement: this.GetString("statement"),
		Description: this.GetString("description"),
		Constraints: this.GetString("constraints"),
		Sample_input: this.GetString("sample_input"),
		Sample_output: this.GetString("sample_output"),
		Type: this.GetString("type"),
		Difficulty: this.GetString("difficulty"),
		Points: points,
		Uid: 1,
	}
	id, noerr := problem.Create()
	if noerr == true {
		pid := strconv.FormatInt(id, 10)
		this.Redirect("/problem/" + pid, 301)
	}

	this.Data["title"] = "Create Problem "

	this.Layout = "layout.tpl"
	this.TplNames = "problem/create.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = ""
	this.LayoutSections["Sidebar"] = ""
}

// Serves the problems list page
func (this *ProblemController) List() {
	problem := models.Problem{}
	problems, _ := problem.GetRecent()
	this.Data["problems"] = problems
	this.Data["title"] = "Home | List "

	this.Layout = "layout.tpl"
	this.TplNames = "problem/list.tpl"
	this.LayoutSections = make(map[string]string)
    this.LayoutSections["HtmlHead"] = ""
    this.LayoutSections["Sidebar"] = "sidebar/showcategories.tpl"
}

// Serves the Problem Page
// To-do : send the name of author as well
func (this *ProblemController) ProblemById() {
	pid := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(pid)
	if err != nil {
        // Redirect to 404
        this.Abort("404")
    }
	p := models.Problem{Pid: id}
	p.GetById()
	this.Data["title"] = p.Statement
	this.Data["problem"] = p

	this.Layout = "layout.tpl"
	this.TplNames = "problem/show.tpl"
	this.LayoutSections = make(map[string]string)
    this.LayoutSections["HtmlHead"] = ""
    this.LayoutSections["Sidebar"] = "sidebar/showsimilar.tpl"
}
