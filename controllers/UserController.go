package controllers

import (
	"OnlineJudge/models"
	"github.com/astaxie/beego"
)

type UserController struct {
	BaseController
}

func (this *UserController) Login() {
	uid := this.GetSession("Uid")
	if uid != nil {
		this.Redirect("/", 302)
	}

	if this.Ctx.Input.Param("0") == "submit" {
		user := models.User{
			Username: this.GetString("username"),
			Password: this.GetString("password"),
		}
		// Handle the flash messages
		err := user.LoginVerify()
		if err != nil {
			flash := beego.NewFlash()
			flash.Error(err.Error())
			flash.Store(&this.Controller)
		}
		if user.Login() == true {
			this.SetSession("Uid", this.GetString("username"))
			user.GetUserInfo()
			this.SetSession("id", user.Uid)

			// store the user ID in the session
			this.Redirect("/", 302)
		}
		//If login failed, flash a relevent message
	}

	this.Data["title"] = "Login"

	this.Layout = "layout.tpl"
	this.TplNames = "user/login.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = ""
	this.LayoutSections["Sidebar"] = ""
	this.LayoutSections["ErrorHead"] = "errorHead.tpl"
}

func (this *UserController) Logout() {
	this.DelSession("Uid")
	this.DelSession("id")
	this.Redirect("/", 302)
}

func (this *UserController) Signup() {
	uid := this.GetSession("Uid")
	if uid != nil {
		this.Redirect("/", 302)
	}

	if this.Ctx.Input.Param("0") != "submit" {
		this.Redirect("/user/login", 302)
	}

	user := models.User{
		Username: this.GetString("username"),
		Password: this.GetString("passkey"),
		Name:     this.GetString("name"),
		College:  this.GetString("college"),
		Email:    this.GetString("email"),
	}
	// All the fields verified, as well checked if username and email are unique
	err := user.SignupVerify()
	if err != nil {
		flash := beego.NewFlash()
		flash.Error(err.Error())
		flash.Store(&this.Controller)
	}
	uid, done := user.Create()

	if done {
		this.SetSession("Uid", this.GetString("username"))
		this.SetSession("id", uid)
		this.Redirect("/", 302)
	}
	this.Redirect("/user/login", 302)
}

// To-do: Show programs solved by the user
func (this *UserController) Show() {
	if this.Ctx.Input.Param("0") == "" {
		this.Abort("404")
	}
	user := models.User{Username: this.Ctx.Input.Param("0")}
	if user.GetByUsername() {
		this.Data["title"] = user.Username
		this.Data["userDetails"] = user

		log := models.Problemlogs{Uid: user.Uid}
		logs, count := log.GetByUid()
		problems := make(map[int]models.Problem)
		if count == 0 {
			this.Data["solvedProblemsExist"] = false
		} else {
			this.Data["solvedProblemsExist"] = true
			for index, element := range logs {
				p := models.Problem{Pid: element.Pid}
				p.GetByPid()
				problems[index] = p
			}
		}

		this.Data["solvedProblems"] = problems

		this.Layout = "layout.tpl"
		this.TplNames = "user/show.tpl"
		this.LayoutSections = make(map[string]string)
		this.LayoutSections["HtmlHead"] = ""
		this.LayoutSections["Sidebar"] = ""
	} else {
		this.Abort("404")
	}
}
