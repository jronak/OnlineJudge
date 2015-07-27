package models

import (
	"github.com/astaxie/beego/orm"
)

func (log *Problemlogs) CommitByPidUid() bool {
	o := orm.NewOrm()
	o.Using("default")
	_, er := o.Insert(log)
	if er == nil {
		return true
	}
	return false
}

func (log *Problemlogs) Update() bool {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Update(log)
	if err == nil {
		return true
	}
	return false
}

func (log *Problemlogs) GetByPidUid() bool {
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(log, "Pid", "Uid")
	if err == nil {
		return true
	}
	return false
}

func (log *Problemlogs) GetRecentByPid() ([]Problemlogs, int64) {
	var logs []Problemlogs
	o := orm.NewOrm()
	o.Using("default")
	count, err := o.QueryTable("problemlogs").OrderBy("-Time").Filter("pid", log.Pid).All(&logs)
	if err == nil {
		return logs, count
	}
	return nil, count
}
