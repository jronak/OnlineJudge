package models

import (
	"github.com/astaxie/beego/orm"
)

func (testcase *Testcases) Add() bool {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Insert(testcase)
	if err == nil {
		return true
	}
	return false
}

func (testcase *Testcases) GetAllByPid() ([]Testcases, int64) {
	var testcases []Testcases
	o := orm.NewOrm()
	o.Using("default")
	count, err := o.QueryTable("testcases").Filter("Pid", testcase.Pid).All(&testcases)
	if err == nil {
		return testcases, count
	}
	return nil, 0
}

func (testcase *Testcases) GetOneByPidTid() (Testcases, bool) {
	var testcaseN Testcases
	o := orm.NewOrm()
	o.Using("default")
	err := o.QueryTable("testcases").Filter("Pid", testcase.Pid).Filter("Tid", testcase.Tid).One(&testcaseN)
	if err == nil {
		return testcaseN, true
	}
	return testcaseN, false
}

func (testcase *Testcases) DeleteAllByPid() (int64, bool) {
	o := orm.NewOrm()
	o.Using("default")
	count, err := o.QueryTable("testcases").Filter("Pid", testcase.Pid).Delete()
	if err == nil {
		return count, true
	}
	return 0, false
}

func (testcase *Testcases) DeleteOneByPidTid() bool {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.QueryTable("testcases").Filter("Pid", testcase.Pid).Filter("Tid", testcase.Tid).Delete()
	if err == nil {
		return true
	}
	return false
}
