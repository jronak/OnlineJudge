package models

import (
	"OnlineJudge/Judge"
)

// Just run code used here. if stdin is nil, then sample input is used
func Exec(pid int, rawCode string, lang string, stdin string) Judge.Code {
	code := Judge.Code{Lang: lang, Stdin: stdin}
	cr := Judge.CRManager{Program: &code, RawCode: &rawCode}
	if stdin == "" {
		problem := new(Problem)
		problem.Pid = pid
		if err := problem.GetSampleIOByPid(); err != true {
			return Judge.Code{}
		}
		code.Stdin = problem.Sample_input
		cr.CROnce()
		if code.Stdout != problem.Sample_output {
			code.RunStatus = false
		}
	} else {
		cr.CROnce()
	}
	return code
}

func ExecBatch(pid int, rawcode string, lang string) []Judge.TestCaseStatus {
	testcase := Testcases{}
	testcase.Pid = pid
	testcases, count := testcase.GetAllByPid()
	j := Judge.TestCaseStatus{}
	j.Comment = "Internal Error"
	j.Success = false
	if count == 0 {
		return []Judge.TestCaseStatus{j}
	}
	testStatus := make([]Judge.TestCaseStatus, count)
	code := Judge.Code{}
	code.Lang = lang
	cr := Judge.CRManager{}
	cr.RawCode = &rawcode
	cr.Program = &code
	cr.CRBatch()
	compileStatus := <-cr.Receive
	if compileStatus == 1 {
		j.Comment = "Compilation Error"
		return []Judge.TestCaseStatus{j}
	}
	for i, test := range testcases {
		cr.Stdin <- test.Input
		cr.Stdout <- test.Output
		testStatus[i] = <-cr.Status
	}
	close(cr.Stdin)
	return testStatus
}
