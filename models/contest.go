package models

import (
	"github.com/astaxie/beego/orm"
	"log"
)

func (this *Contest) Create() (int64, bool) {
	o := orm.NewOrm()
	o.Using("default")
	id, err := o.Insert(this)
	if err != nil {
		log.Println("Contest Model: ", err)
		return -1, false
	}
	return id, true
}

func (this *Contest) Edit() bool {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Update(this)
	if err != nil {
		log.Println("Contest Model: ", err)
		return false
	}
	return true
}

func (this *Contest) GetByName() bool {
	o := orm.NewOrm()
	o.Using("default")
	err := o.QueryTable("contest").Filter("name", this.Name).One(this)
	if err != nil {
		return false
	}
	return true
}

func (this *Contest) Delete() bool {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Delete(this)
	if err != nil {
		log.Println("Contest Model: ", err)
		return false
	}
	return true
}
