package controllers

import (
	"OnlineJudge/models"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

/*
	if session is registeried, redirected to root(/)
	else login page displayed
	Key = unique for keeping away machine; if aint worth, ditch
	error = Using error to display on test login
*/
func (this *LoginController) Get() {
	sess := this.StartSession()
	uid := sess.Get("Uid")
	fmt.Println(uid)
	if uid != nil {
		this.Redirect("/", 302)
	}
	this.Data["error"] = ""
	this.Data["key"] = ""
	this.TplNames = "login.html"
}

//Username and login to be received via header
func (this *LoginController) Post() {
	username := this.GetString("username")
	password := this.GetString("password")
	ubool := models.CheckUserName(username)
	pbool := models.CheckPassword(password)
	if ubool && pbool {
		user := &models.User{Username: username, Password: password}
		if user.Login() {
			fmt.Println(user.Uid)
			ses := this.StartSession()
			ses.Set("Uid", user.Uid)
			ses.Set("Username", username)
		} else {
			this.Data["error"] = "Username or password is incorrect"
			this.TplNames = "login.html"
			return
		}
	} else {
		this.Data["error"] = "Username or password is invalid"
		this.TplNames = "login.html"
		return
	}
	this.Redirect("/", 302)
}
