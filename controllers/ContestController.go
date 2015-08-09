package controllers

import (
	"OnlineJudge/models"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"
)

type ContestController struct {
	BaseController
}

func (this *ContestController) isAdmin() bool {
	if !this.isLoggedIn() {
		this.Redirect("/user/login", 302)
		return false
	}
	name := this.GetSession("Uid")
	if name.(string) != "admin" {
		this.Redirect("/", 302)
		return false
	}
	return true
}

func (this *ContestController) isLoggedIn() bool {
	if this.GetSession("id") != nil {
		return true
	}
	return false
}

//Serve the create contest page
// /contest/create
func (this *ContestController) CreateContest() {
	if !this.isAdmin() {
		return
	}

}

// /contest/save
func (this *ContestController) SaveContest() {
	if !this.isAdmin() {
		return
	}
	contest := models.Contest{}
	contest.Name = this.GetString("name")
	contest.Description = this.GetString("description")
	_, _ = contest.Create()
	this.Redirect("/contest/"+contest.Name+"/addProblem", 302)
}

// /contest/:name/addproblem
func (this *ContestController) AddProblem() {
	if !this.isAdmin() {
		return
	}
	//Serve the page
}

// /contest/:name/saveproblem
func (this *ContestController) saveProblem() {
	if !this.isAdmin() {
		return
	}
	contestName := this.Ctx.Input.Param(":name")
	id := this.GetSession("id")
	user := models.User{}
	user.Uid = id.(int)
	points, _ := strconv.Atoi(this.GetString("points"))
	//remove replace foe newlines
	sampleInput := strings.Replace(this.GetString("sample_input"), "\r", "", -1)
	sampleOutput := strings.Replace(this.GetString("sample_output"), "\r", "", -1)
	problem := models.Problem{
		Statement:   this.GetString("statement"),
		Description: this.GetString("description"),
		Constraints: this.GetString("constraints"),
		//Sample_input:  this.GetString("sample_input"),
		//Sample_output: this.GetString("sample_output"),
		Difficulty: this.GetString("difficulty"),
		Points:     points,
		Uid:        id.(int),
	}
	problem.Sample_output = sampleOutput
	problem.Sample_input = sampleInput
	problem.Type = "contest" + contestName
	pid, status := problem.Create()
	if status {
		id := strconv.Itoa(pid)
		this.Redirect("/contest/"+contestName+"/"+id, 302)
	}
	// handle the failure
}

// /contest/:name/:pid
func (this *ContestController) GetProblem() {
	pid := this.Ctx.Input.Param(":id")
	contestName := this.Ctx.Input.Param(":name")
	id, err := strconv.Atoi(pid)
	if err != nil {
		// Redirect to 404
		this.Abort("404")
	}
	p := models.Problem{Pid: id}
	p.GetById()
	check := strings.Contains(p.Type, contestName)
	if !check {
		this.Redirect("/contest/"+contestName, 302)
		return
	}
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

	this.Data["title"] = p.Statement
	this.Data["problem"] = p
	this.Data["Author"] = contestName
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
}

// /contest/:name/:pid/submit
func (this *ContestController) submit() {
	if !this.isLoggedIn() {
		this.Redirect("/user/login", 302)
		return
	}
	contestName := this.Ctx.Input.Param(":name")
	pid, err := strconv.Atoi(this.Ctx.Input.Param(":pid"))
	if err != nil {
		log.Println(err)
		this.Redirect("/", 302)
		return
	}
	contest := models.Contest{}
	contest.Name = contestName
	if err := contest.GetByName(); !err {
		this.Redirect("/", 302)
		return
	}
	currTime := time.Now()
	if contest.EndTime.Before(currTime) {
		this.Redirect("/contest/"+contestName, 302)
		return
	}
	uid := this.GetSession("id")
	code := this.GetString("code")
	lang := this.GetString("language")

	problemLog := models.Problemlogs{}
	problemLog.Uid = uid.(int)
	problemLog.Pid = pid
	problemLog.GetByPidUid()

	output := models.SubmitUpdateScore(uid.(int), pid, code, lang)

	contestLog := models.Contestlogs{}
	contestLog.Uid = uid.(int)
	contestLog.Cid = contest.Id
	status := contestLog.GetByUidCid()
	if !status {
		contestLog.Points = output.Score
		contestLog.Add()
	} else {
		if output.Score > problemLog.Points {
			contestLog.Points += output.Score - problemLog.Points
			contestLog.Update()
		}
	}

	js, _ := json.Marshal(output)
	this.Data["json"] = string(js)
	this.ServeJson()

}

//Serve the contest page
// /contest/:name
func (this *ContestController) Contest() {

	contestName := this.Ctx.Input.Param(":name")
	contest := models.Contest{}
	contest.Name = contestName
	if err := contest.GetByName(); !err {
		this.Redirect("/", 302)
		return
	}
	currTime := time.Now()
	if contest.EndTime.Before(currTime) {
		this.Data["elapsed"] = true
	}
}
