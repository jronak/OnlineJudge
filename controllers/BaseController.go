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

	this.Data["logged"] = false
	if uname != nil {
		user := models.User{ Username: uname.(string) }
		this.Data["user"] = user
		this.Data["login"] = uname.(string)
		this.Data["logged"] = true
	}
}