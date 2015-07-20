package controllers

import (
	"OnlineJudge/models"
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
		if user.Login() == true {
			this.SetSession("Uid", this.GetString("username"))
			this.Redirect("/", 302)
		}
	}

	this.Data["title"] = "Login"

	this.Layout = "layout.tpl"
	this.TplNames = "user/login.tpl"
	this.LayoutSections = make(map[string]string)
    this.LayoutSections["HtmlHead"] = ""
    this.LayoutSections["Sidebar"] = ""
}

func (this *UserController) Logout() {
	this.DelSession("Uid")
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
		Name: this.GetString("name"),
		College: this.GetString("college"),
		Email: this.GetString("email"),
	}

	uid, done := user.Create()

	if done {
		this.SetSession("Uid", this.GetString("username"))
		this.Redirect("/", 302)
	}
	this.Redirect("/user/login", 302)
}
