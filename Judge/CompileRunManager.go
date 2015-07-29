package Judge

import (
	"io/ioutil"
	"log"
)

const (
	defaultPath         = "/tmp/OnlineJ/Default/"
	javaPaths           = "/tmp/OnlineJ/JavaP/"
	rootPath            = "/tmp/OnlineJ"
	fileDirPermissions  = 0755
	javaPathsBufferSize = 1
	nameBufferSize      = 10
	dirSuffixLen        = 2
	fileNameLen         = 6
	crOnceWorkSize      = 2
	crOnceChanBufSize   = 2 * crOnceWorkSize
	deleteBufferSize    = 2 * crOnceChanBufSize
)

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

var (
	crChannelOnce     chan *CRManager
	crChannelBatch    chan *CRManager
	supportedLangs    []string
	codeExtensionsMap map[string]string
)

func (cr *CRManager) isCompilableLang() bool {
	_, val := CompileCommandMap[cr.Program.Lang]
	return val
}

func (cr *CRManager) createFile() {
	path := cr.Program.path + cr.Program.name + codeExtensionsMap[cr.Program.Lang]
	err := ioutil.WriteFile(path, []byte(*cr.RawCode), fileDirPermissions)
	CheckError(err)
}

// Worker for Compile once
func crOnceWorker(channel chan *CRManager) {
	for cr := range channel {
		run := true
		cr.createFile()
		if cr.isCompilableLang() {
			cr.Program.CompilationManager()
			if cr.Program.CompilationStatus == false {
				cr.receive <- 1
				run = false
			}
		}
		if run {
			cr.Program.RunManager()
			cr.receive <- 1
		}
	}
}

// Compiler and run once
func (cr *CRManager) CROnce() {
	cr.receive = make(chan int)
	if cr.isCompilableLang() {
		cr.Program.execTimeLimit = defaultExecTimeout
	} else {
		cr.Program.execTimeLimit = defaultExecTimeoutScript
	}
	crChannelOnce <- cr
	<-cr.receive
}

func crBatchWorker(channel chan *CRManager) {
	for cr := range channel {
		run := true
		cr.createFile()
		if cr.isCompilableLang() {
			cr.Program.CompilationManager()
			if cr.Program.CompilationStatus == false {
				cr.receive <- 1
				run = false
			} else {
				cr.receive <- 0
			}
		} else {
			cr.receive <- 0
		}
		if run {
			cr.RunBatchManager()
		}
	}
}

func (cr *CRManager) CRBatch() {
	cr.receive = make(chan int, 1)
	cr.status = make(chan TestCaseStatus, 1)
	cr.stdin = make(chan string, 1)
	cr.stdout = make(chan string, 1)
	crChannelBatch <- cr
	testStatus := make([]TestCaseStatus, len(cr.TestInput))
	cr.CRBatch()
	compileStatus := <-cr.receive
	if compileStatus == 1 {
		j := TestCaseStatus{}
		j.Comment = "Compilation Error"
		j.Success = false
		cr.TestCaseOutput = []TestCaseStatus{j}
		return
	}
	for i, stdin := range cr.TestInput {
		cr.stdin <- stdin
		cr.stdout <- cr.TestOutput[i]
		testStatus[i] = <-cr.status
	}
	cr.TestCaseOutput = testStatus
	close(cr.stdin)
}

func (cr *CRManager) CR() {
	if cr.Isbatch == true {
		cr.CRBatch()
	} else {
		cr.CROnce()
	}
}

func init() {
	crChannelOnce = make(chan *CRManager, crOnceChanBufSize)
	crChannelBatch = make(chan *CRManager, crOnceChanBufSize)
	supportedLangs = []string{"C", "C++", "Java", "Python3", "Python2", "Go"}
	extensions := []string{".c", ".cpp", ".java", ".py", ".py", ".go"}
	codeExtensionsMap = make(map[string]string)
	for iter, extn := range supportedLangs {
		codeExtensionsMap[extn] = extensions[iter]
	}
	for iter := 0; iter < crOnceWorkSize; iter++ {
		go crOnceWorker(crChannelOnce)
		go crBatchWorker(crChannelBatch)
	}
	go log.Println("CRManager Init : Normal")
}
