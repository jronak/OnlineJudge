package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func (problem *Problem) Create() (int64, bool) {
	o := orm.NewOrm()
	o.Using("default")
	id, err := o.Insert(problem)
	if err == nil {
		return id, true
	}
	beego.Error(err)
	return 0, false
}

func (problem *Problem) Update() bool {
	o := orm.NewOrm()
	o.Using("default")
	_, b := o.Update(problem)
	if b == nil {
		return true
	}
	return false
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

func (problem *Problem) GetRecent() ([]Problem, int64) {
	var problems []Problem
	o := orm.NewOrm()
	o.Using("default")
	count, err := o.QueryTable("problem").OrderBy("-Created_at").Limit(10).All(&problems)
	if err == nil {
		return problems, count
	}
	return nil, count
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

func (problem *Problem) GetByStatement() bool {
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(problem)
	if err == nil {
		return true
	}
	return false
}

func (problem *Problem) GetById() bool {
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(problem)
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

func (problem *Problem) GetTypes() (*orm.ParamsList, int64) {
	list := new(orm.ParamsList)
	o := orm.NewOrm()
	o.Using("default")
	num, _ := o.Raw("SELECT DISTINCT type from problem").ValuesFlat(list)
	return list, num
}

func (problem *Problem) GetByType(page int) ([]Problem, int64) {
	var problems []Problem
	o := orm.NewOrm()
	o.Using("default")
	count, err := o.QueryTable("problem").Filter("type", problem.Type).Offset((page-1)*10).Limit(10).All(&problems)
	if err == nil {
		return problems, count
	}
	return nil, count
}

func (problem *Problem) GetSampleIOByPid() bool {
	o := orm.NewOrm()
	o.Using("default")
	err := o.QueryTable("problem").Filter("pid", problem.Pid).One(problem, "sample_input", "sample_output")
	if err == nil {
		return true
	}
	return false
}
