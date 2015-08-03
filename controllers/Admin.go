package controllers

import (
	"OnlineJudge/models"
	"encoding/json"
	"strconv"
)

type AdminController struct {
	BaseController
}

func (this *AdminController) isAdmin() bool {
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

func (this *AdminController) isLoggedIn() bool {
	if this.GetSession("id") != nil {
		return true
	}
	return false
}

func (this *AdminController) ShowEditors() {
	if !this.isAdmin() {
		return
	}
	user := models.User{}
	users := user.GetEditors()
	bytes, _ := json.Marshal(users)
	this.Data["json"] = string(bytes)
	this.ServeJson()
}

// /admin/makeEditor/:uid
func (this *AdminController) MakeEditor() {
	if !this.isAdmin() {
		return
	}
	uid := this.Ctx.Input.Param("uid")
	user := models.User{}
	id, _ := strconv.Atoi(uid)
	user.Uid = id
	status := user.MakeEditor()
	this.Data["status"] = status
}

// /admin/revokeEditor/:uid
func (this *AdminController) RevokeEditor() {
	if !this.isAdmin() {
		return
	}
	uid := this.Ctx.Input.Param("uid")
	user := models.User{}
	id, _ := strconv.Atoi(uid)
	user.Uid = id
	status := user.RevokeEditor()
	this.Data["status"] = status
}

// /admin/search/name/:name
func (this *AdminController) SearchName() {
	if !this.isAdmin() {
		return
	}
	name := this.Ctx.Input.Param(":name")
	user := models.User{}
	user.Name = name
	users, _ := user.SearchByName()
	bytes, _ := json.Marshal(users)
	this.Data["json"] = string(bytes)
	this.ServeJson()
}

// /admin/search/name/:uid
func (this *AdminController) DeleteUser() {
	if !this.isAdmin() {
		return
	}

	uid := this.Ctx.Input.Param("uid")
	user := models.User{}
	id, _ := strconv.Atoi(uid)
	user.Uid = id
	status := user.Delete()
	this.Data["status"] = status
}
