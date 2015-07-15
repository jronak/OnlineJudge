package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Uid       int `orm:"pk"`
	Username  string
	Password  string `json:"-"`
	Name      string
	College   string
	Email     string
	Score     int
	Rank      int
	Is_editor int `json:"-"`
}

type Problem struct {
	Pid                  int `orm:"pk"`
	Uid                  int
	Statement            string
	Description          string
	Constraints          string
	Sample_input         string
	Sample_output        string
	Solution_description string `json:"-"`
	Solution_code        string `json:"-"`
	Type                 string
	Difficulty           string
	Created_at           time.Time
	Points               int
	Solve_count          int
}

type Testcases struct {
	Id      int `orm:"pk"`
	Pid     int
	Tid     int
	Input   string
	Output  string
	Timeout int
}

type Problemlogs struct {
	Id     int `orm:"pk"`
	Pid    int
	Uid    int
	Solved int
	Points int
	Time   time.Time
}

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "ronak:ronak@/OnlineJudge")
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Problem))
	orm.RegisterModel(new(Problemlogs))
	orm.RegisterModel(new(Testcases))
}
