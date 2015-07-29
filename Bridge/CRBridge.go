package Bridge

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	containerPath     = "/var/lib/lxc/"
	containerPrefix   = "judge"
	dirPrefix         = "/rootfs/home/ubuntu/d"
	problemName       = "problem.json"
	solutionName      = "solution.json"
	regularContainers = 1
	contestContainers = 0
	dirsCount         = 1
)

var (
	dirChannel chan string
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
	dir := <-dirChannel
	bytes, err := json.Marshal(cr)
	if err != nil {
		loghelper(err)
		return
	}
	file, err := os.Create(dir + problemName)
	if err != nil {
		loghelper(err)
		return
	}
	_, _ = file.WriteString(string(bytes))
	i := true
	for i == true {
		if solution, _ := os.OpenFile(dir+solutionName, os.O_RDWR, 0755); solution != nil {
			bytes, err := ioutil.ReadAll(solution)
			err = json.Unmarshal(bytes, cr)
			if err != nil {
				loghelper(err)
			}
			break
		}
		time.Sleep(time.Second * 1)
	}
	go func() {
		os.RemoveAll(dir)
		dirChannel <- dir
	}()
}

//TO_DO change the container name from test to ubuntu
func init() {
	dirChannel = make(chan string, dirsCount)
	for containers := 1; containers <= regularContainers; containers++ {
		containerSuffix := strconv.Itoa(containers)
		for i := 1; i <= dirsCount; i++ {
			dir := containerPath + containerPrefix + containerSuffix + dirPrefix + strconv.Itoa(i) + "/"
			fmt.Println("Dir", dir)
			dirChannel <- dir
		}
	}
}
