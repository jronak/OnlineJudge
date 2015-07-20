package models

import (
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func (user *User) IsUsernameUnique() bool {
	o := orm.NewOrm()
	o.Using("default")
	return !o.QueryTable("user").Filter("username", user.Username).Exist()
}

func (user *User) IsEmailUnique() bool {
	o := orm.NewOrm()
	o.Using("default")
	return !o.QueryTable("user").Filter("email", user.Password).Exist()
}

func (user *User) Create() (int64, bool) {
	o := orm.NewOrm()
	o.Using("default")
	uid, err := o.Insert(user)
	if err == nil {
		user.Password = ""
		return uid, true
	}
	return 0, false
}

func (user *User) Login() bool {
	o := orm.NewOrm()
	o.Using("default")
	err := o.QueryTable("user").Filter("username", user.Username).Filter("password", user.Password).One(user, "uid")
	if err == nil {
		return true
	}
	return false
}

func (user *User) UpdateCollege(college string) bool {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.QueryTable("user").Filter("uid", user.Uid).Update(orm.Params{"college": college})
	if err == nil {
		return true
	}
	return false
}

func (user *User) MakeEditor() bool {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.QueryTable("user").Filter("uid", user.Uid).Update(orm.Params{"is_editor": 1})
	if err == nil {
		return true
	}
	return false
}

func (user *User) RevokeEditor() bool {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.QueryTable("user").Filter("uid", user.Uid).Update(orm.Params{"is_editor": 0})
	if err == nil {
		return true
	}
	return false
}

func (user *User) ChangePassword(password string) bool {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.QueryTable("user").Filter("uid", user.Uid).Update(orm.Params{"password": password})
	if err == nil {
		return true
	}
	return false
}

func (user *User) GetUserInfo() bool {
	o := orm.NewOrm()
	o.Using("default")
	err := o.QueryTable("user").Filter("uid", user.Uid).One(user, "uid", "username", "name",
		"college", "email", "score", "rank")
	if err == nil {
		return true
	}
	return false
}

func (user *User) Delete() bool {
	o := orm.NewOrm()
	o.Using("default")
	_, err := o.Delete(user)
	if err == nil {
		return true
	}
	return false

}

func (user *User) IsEditor() bool {
	o := orm.NewOrm()
	o.Using("default")
	err := o.Read(user)
	user.Password = ""
	if err == nil {
		if user.Is_editor == 1 {
			return true
		}
	}
	return false
}

func (user *User) AddScore(score int) bool {
	o := orm.NewOrm()
	o.Using("default")
	if err := o.Read(user); err == nil {
		user.Score += score
		if _, err := o.Update(user); err == nil {
			return true
		}
	}
	return false
}

func (user *User) UpdateRank(rank int) bool {
	o := orm.NewOrm()
	o.Using("default")
	if err := o.Read(user); err == nil {
		user.Rank = rank
		if _, err := o.Update(user); err == nil {
			return true
		}
	}
	return false
}

func (user *User) Get() bool {
	o := orm.NewOrm()
	o.Using("default")
	if err := o.Read(user); err == nil {
		return true
	}
	return false
}
