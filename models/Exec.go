package models

import (
	"OnlineJudge/Bridge"
)

// Just run code used here. if stdin is nil, then sample input is used
func Exec(pid int, rawCode string, lang string, stdin string) Bridge.Code {
	code := Bridge.Code{Lang: lang, Stdin: stdin}
	cr := Bridge.CRManager{Program: &code, RawCode: &rawCode}
	if stdin == "" {
		problem := new(Problem)
		problem.Pid = pid
		if err := problem.GetSampleIOByPid(); err != true {
			return Bridge.Code{}
		}
		code.Stdin = problem.Sample_input
		Bridge.CompileExec(&cr)
		if code.Stdout != problem.Sample_output {
			code.RunStatus = false
		}
	} else {
		Bridge.CompileExec(&cr)
	}
	return code
}

func ExecBatch(pid int, rawcode string, lang string) []Bridge.TestCaseStatus {
	testcase := Testcases{}
	testcase.Pid = pid
	testcases, count := testcase.GetAllByPid()
	j := Bridge.TestCaseStatus{}
	j.Comment = "Internal Error"
	j.Success = false
	if count == 0 {
		return []Bridge.TestCaseStatus{j}
	}
	code := Bridge.Code{}
	code.Lang = lang
	cr := Bridge.CRManager{}
	cr.RawCode = &rawcode
	cr.Program = &code
	cr.Isbatch = true
	testInput := make([]string, count)
	testOutput := make([]string, count)
	for i, test := range testcases {
		testInput[i] = test.Input
		testOutput[i] = test.Output

	}
	cr.TestInput = testInput
	cr.TestOutput = testOutput
	Bridge.CompileExec(&cr)
	return cr.TestCaseOutput
}
