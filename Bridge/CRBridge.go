package Bridge

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	containerPrefix = "judge"
	containerCount  = 1
)

var (
	ipChan chan string
)

type Code struct {
	path              string
	name              string
	Lang              string
	Stderr            string
	Stdout            string
	Stdin             string
	CompilationStatus bool
	RunStatus         bool
	execTimeLimit     time.Duration
	ExecTime          float64
}

type TestCaseStatus struct {
	Success  bool
	Comment  string
	ExecTime float64
}

type CRManager struct {
	Program        *Code
	TestCaseOutput []TestCaseStatus
	TestInput      []string
	TestOutput     []string
	RawCode        *string
	Isbatch        bool
	receive        chan int
	stdin          chan string
	stdout         chan string
	status         chan TestCaseStatus
}

func loghelper(err error) {
	log.Println("Bridge Error: ", err)
}

func CompileExec(cr *CRManager) {
	ip := <-ipChan
	bytes, _ := json.Marshal(cr)
	resp, _ := http.PostForm(ip, url.Values{"json": {string(bytes)}})
	solu := resp.Header.Get("json")
	resp.Body.Close()
	json.Unmarshal([]byte(solu), cr)
	ipChan <- ip
}

//TO_DO change the container name from test to ubuntu
func init() {
	ipChan = make(chan string, containerCount)
	containerIP := "http://10.0.3.136:8080"
	ipChan <- containerIP
}
