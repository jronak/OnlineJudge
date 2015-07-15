package models

import (
	"github.com/astaxie/beego/orm"
)

func (problem *Problem) Create() (int64, bool) {
	o := orm.NewOrm()
	o.Using("default")
	id, err := o.Insert(problem)
	if err == nil {
		return id, true
	}
	return 0, false
}

func (problem *Problem) GetByPid() bool {
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(problem)
	if err == nil {
		return true
	}
	return false
}

func (problem *Problem) DeleteByPid() bool {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Delete(problem)
	if err == nil {
		return true
	}
	return false
}

func (problem *Problem) GetByName() bool {
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(problem, "statement")
	if err == nil {
		return true
	}
	return false
}

func (problem *Problem) GetByUid() ([]Problem, int64) {
	var problems []Problem
	o := orm.NewOrm()
	o.Using("default")
	count, err := o.QueryTable("problem").Filter("uid", problem.Uid).All(&problems, "pid", "statement", "type",
		"difficulty", "points", "solve_count")
	if err == nil {
		return problems, count
	}
	return nil, count

}

func (problem *Problem) GetByType() ([]Problem, int64) {
	var problems []Problem
	o := orm.NewOrm()
	o.Using("default")
	count, err := o.QueryTable("problem").Filter("type", problem.Type).All(&problems, "pid", "statement", "type",
		"difficulty", "points", "solve_count")
	if err == nil {
		return problems, count
	}
	return nil, count
}
