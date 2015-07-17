package models

import (
	"OnlineJudge/Judge"
	"time"
)

type SubmitResponse struct {
	Score  int
	Status []Judge.TestCaseStatus
}

func SubmitUpdateScore(uid int, pid int, rawcode string, lang string) SubmitResponse {
	testCaseStatus := ExecBatch(pid, rawcode, lang)
	submitResponse := SubmitResponse{}
	submitResponse.Status = testCaseStatus
	problem := Problem{}
	problem.Pid = pid
	_ = problem.GetByPid()
	problemlog := Problemlogs{}
	problemlog.Pid = pid
	problemlog.Uid = uid
	b := problemlog.GetByPidUid()
	score := computeScore(problem.Points, testCaseStatus)
	if score == 0 {
		return submitResponse
	}
	if b == false {
		problemlog.Uid = uid
		problemlog.Pid = pid
		problemlog.Solved = 1
		problemlog.Points = score
		problemlog.Time = time.Now()
		problemlog.CommitByPidUid()
	} else {
		if problemlog.Points < score {
			problemlog.Points = score
			problemlog.Update()
		}
	}
	submitResponse.Score = score
	return submitResponse
}

func computeScore(maxScore int, testCaseStatus []Judge.TestCaseStatus) int {
	casePassed := 0
	for _, status := range testCaseStatus {
		if status.Success == true {
			casePassed++
		}
	}
	return (casePassed * maxScore) / len(testCaseStatus)
}
