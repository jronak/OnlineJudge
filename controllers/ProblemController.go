package controllers

import (
	"OnlineJudge/models"
	"encoding/json"
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
	this.Data["types"], _ = problem.GetTypes()

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

	//Redirect if user doesnt hold editor rights
	id := this.GetSession("id")
	user := models.User{}
	user.Uid = id.(int)
	if !user.IsEditor() {
		this.Redirect("/", 302)
		return
	}

	points, _ := strconv.Atoi(this.GetString("points"))
	//remove replace foe newlines
	problem := models.Problem{
		Statement:     this.GetString("statement"),
		Description:   this.GetString("description"),
		Constraints:   this.GetString("constraints"),
		Sample_input:  this.GetString("sample_input"),
		Sample_output: this.GetString("sample_output"),
		Type:          this.GetString("type"),
		Difficulty:    this.GetString("difficulty"),
		Points:        points,
		Uid:           id.(int),
	}
	id, noerr := problem.Create()
	pid := strconv.Itoa(id.(int))
	if noerr == true {
		this.Redirect("/problem/"+pid, 302)
	}

	this.Data["title"] = "Create Problem "

	this.Layout = "layout.tpl"
	this.TplNames = "problem/create.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = ""
	this.LayoutSections["Sidebar"] = ""
}

func (this *ProblemController) AddTestCase() {
	if !this.isLoggedIn() {
		this.Redirect("/user/login", 302)
		return
	}

	//Redirect if user doesnt hold editor rights
	uid := this.GetSession("id")
	user := models.User{Uid: uid.(int)}
	if !user.IsEditor() {
		this.Redirect("/", 302)
		return
	}

	pid := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(pid)
	problem := models.Problem{Pid: id}
	problem.GetByPid()
	this.Data["problem"] = problem

	testcases := models.Testcases{Pid: id}
	cases, _ := testcases.GetAllByPid()

	this.Data["title"] = "Add Test Case"
	this.Data["cases"] = cases

	this.Layout = "layout.tpl"
	this.TplNames = "problem/addtest.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = ""
	this.LayoutSections["Sidebar"] = ""

}

func (this *ProblemController) SaveTestCase() {
	if !this.isLoggedIn() {
		this.Redirect("/user/login", 302)
		return
	}

	//Redirect if user doesnt hold editor rights
	uid := this.GetSession("id")
	user := models.User{Uid: uid.(int)}
	if !user.IsEditor() {
		this.Redirect("/", 302)
		return
	}

	pid := this.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(pid)

	timeout, _ := strconv.Atoi(this.GetString("timeout"))
	//remove string replace
	testcase := models.Testcases{
		Pid:     id,
		Input:   this.GetString("input"),
		Output:  this.GetString("output"),
		Timeout: timeout,
	}

	done := testcase.Create()
	if done == true {
		this.Redirect("/problem/"+pid, 302)
	}
	this.Redirect("/problem/"+pid+"/addtest", 302)
}

// Serves the problems list page
func (this *ProblemController) List() {
	problem := models.Problem{}
	problems, _ := problem.GetRecent()
	this.Data["problems"] = problems
	this.Data["title"] = "Home | List "
	this.Data["types"], _ = problem.GetTypes()

	this.Layout = "layout.tpl"
	this.TplNames = "problem/list.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = ""
	this.LayoutSections["Sidebar"] = "sidebar/showcategories.tpl"
}

// Serves the Problem Page
// To-do: Show recently solved users and their language on sidebar
// To-do: Later, add least execution time log on sidebar
func (this *ProblemController) ProblemById() {
	pid := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(pid)
	if err != nil {
		// Redirect to 404
		this.Abort("404")
	}
	p := models.Problem{Pid: id}
	p.GetById()

	log := models.Problemlogs{Pid: id}
	logs, count := log.GetRecentByPid()
	users := make(map[int]models.User)
	if count == 0 {
		this.Data["recentlySolvedUsersExist"] = false
	} else {
		this.Data["recentlySolvedUsersExist"] = true
		for index, element := range logs {
			u := models.User{Uid: element.Uid}
			u.GetUserInfo()
			users[index] = u
		}
	}

	//Author added
	user := models.User{}
	user.Uid = p.Uid
	user.GetUserInfo()
	this.Data["title"] = p.Statement
	this.Data["problem"] = p
	this.Data["Author"] = user.Username
	this.Data["recentlySolvedUsers"] = users

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
	this.LayoutSections["Sidebar"] = "sidebar/recently_solved_by.tpl"
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
	code := this.GetString("code")
	lang := this.GetString("language")
	output := models.SubmitUpdateScore(uid.(int), pid, code, lang)
	js, _ := json.Marshal(output)
	this.Data["json"] = string(js)
	this.ServeJson()

}

func (this *ProblemController) RunCode() {
	if !this.isLoggedIn() {
		this.Redirect("/user/login", 302)
		return
	}

	//uid := this.GetSession("id")
	pid, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	problem := models.Problem{Pid: pid}
	problem.GetByPid()
	code := this.GetString("code")
	lang := this.GetString("language")
	stdin := this.GetString("stdin")
	output := models.Exec(pid, code, lang, stdin)
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
