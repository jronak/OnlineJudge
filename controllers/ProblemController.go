package controllers

import (
	"OnlineJudge/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type ProblemTypes struct {
	Count      int64
	Categories *orm.ParamsList
}

type ProblemController struct {
	BaseController
}

// Using list template here as well
// To do : serve in pages, per page 10 problems - done
func (this *ProblemController) ProblemByCategory() {
	problemType := this.Ctx.Input.Param(":type")
	page, _ := strconv.Atoi(this.Ctx.Input.Param(":page"))
	problem := models.Problem{Type: problemType}
	problems, count := problem.GetByType(page)
	if count == 0 {
		this.Redirect("/", 302)
		return
	}
	this.Data["problems"] = problems
	this.Data["title"] = "Home | List "

	this.Layout = "layout.tpl"
	this.TplNames = "problem/list.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = ""
	this.LayoutSections["Sidebar"] = "sidebar/showcategories.tpl"
}

// Create Page
func (this *ProblemController) Create() {

	// If not logged redirect to login
	if !this.isLoggedIn() {
		this.Redirect("/user/login", 302)
		return
	}

	//Redirect if user doesnt hold editor rights
	id := this.GetSession("id")
	user := models.User{}
	user.Uid = id.(int)
	if !user.IsEditor() {
		this.Redirect("/", 302)
		return
	}

	this.Data["title"] = "Create Problem "
	this.Layout = "layout.tpl"
	this.TplNames = "problem/create.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = ""
	this.LayoutSections["Sidebar"] = ""
}

// Save Problem
// To-do: Clean info before save - Ambigous
// To-do: Check login and user previlages - Done
func (this *ProblemController) SaveProblem() {

	if !this.isLoggedIn() {
		this.Redirect("/user/login", 302)
		return
	}

	points, _ := strconv.Atoi(this.GetString("points"))
	problem := models.Problem{
		Statement:     this.GetString("statement"),
		Description:   this.GetString("description"),
		Constraints:   this.GetString("constraints"),
		Sample_input:  this.GetString("sample_input"),
		Sample_output: this.GetString("sample_output"),
		Type:          this.GetString("type"),
		Difficulty:    this.GetString("difficulty"),
		Points:        points,
		Uid:           1,
	}
	id, noerr := problem.Create()
	if noerr == true {
		pid := strconv.FormatInt(id, 10)
		this.Redirect("/problem/"+pid, 301)
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
// To-do : send the name of author as well - Done
func (this *ProblemController) ProblemById() {
	pid := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		// Redirect to 404
		this.Abort("404")
	}
	p := models.Problem{Pid: id}
	p.GetById()

	//Author added
	user := models.User{}
	user.Uid = p.Uid
	user.GetUserInfo()
	this.Data["title"] = p.Statement
	this.Data["problem"] = p
	this.Data["Author"] = user.Username

	// Handle problem log of a user
	if this.isLoggedIn() {
		problemLog := models.Problemlogs{}
		problemLog.Pid = p.Pid
		problemLog.Uid = p.Uid
		if problemLog.GetByPidUid() {
			this.Data["userScore"] = problemLog.Points
			this.Data["solvedCount"] = problemLog.Solved
		}
	}

	this.Layout = "layout.tpl"
	this.TplNames = "problem/show.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "problem/submit_head.tpl"
	this.LayoutSections["Sidebar"] = "sidebar/showsimilar.tpl"
}

// Format of submission
// Status - crashes on submission, code and lang are empty
func (this *ProblemController) SaveSubmission() {
	if !this.isLoggedIn() {
		this.Redirect("/user/login", 302)
		return
	}
	uid := this.GetSession("id")
	pid, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	code := this.Data["code"]
	lang := this.Data["language"]
	fmt.Println(pid, uid, code, lang)
	output := models.SubmitUpdateScore(uid.(int), pid, code.(string), lang.(string))
	js, _ := json.Marshal(output)
	this.Data["json"] = string(js)
	this.ServeJson()

}

func (this *ProblemController) isLoggedIn() bool {
	if this.GetSession("id") != nil {
		return true
	}
	return false
}
