package models

import (
	"github.com/astaxie/beego/orm"
)

func (score *Contestlogs) Add() bool {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Insert(score)
	if err != nil {
		return false
	}
	return true
}

func (score *Contestlogs) GetByUidCid() bool {
	o := orm.NewOrm()
	o.Using("default")
	err := o.QueryTable("contestlogs").Filter("uid", score.Uid).Filter("cid", score.Cid).One(score)
	if err != nil {
		return false
	}
	return true
}

func (score *Contestlogs) Update() bool {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Update(score)
	if err != nil {
		return false
	}
	return true
}

func (score *Contestlogs) GetByCid() (*[]Contestlogs, bool) {
	logs := new([]Contestlogs)
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.QueryTable("contestlogs").Filter("cid", score.Cid).OrderBy("-points").All(logs)
	if err != nil {
		return nil, false
	}
	return logs, true
}
