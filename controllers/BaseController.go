package controllers

import (
	"github.com/astaxie/beego"
	"OnlineJudge/models"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	uname := this.GetSession("Uid")
	uid := this.GetSession("id")

	this.Data["logged"] = false
	if uname != nil {
		user := models.User{ Uid: uid.(int), Username: uname.(string) }
		_ = user.GetUserInfo()
		this.Data["isEditor"] = false
		if user.IsEditor() {
			this.Data["isEditor"] = true
		}
		this.Data["user"] = user
		this.Data["login"] = uname.(string)
		this.Data["logged"] = true
	}
}