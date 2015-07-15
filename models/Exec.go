package models

import (
	"OnlineJudge/Judge"
)

func Exec(rawCode string, lang string, stdin string) *Judge.Code {
	code := Judge.Code{Lang: lang, Stdin: stdin}
	cr := Judge.CRManager{Program: &code, RawCode: &rawCode}
	cr.CROnce()
	return &code
}
