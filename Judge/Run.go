package Judge

import (
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"
)

type RunResponse struct {
	code        *Code
	runResponse chan int
}

type TestCaseStatus struct {
	Success  bool
	Comment  string
	ExecTime float64
}

const (
	defaultExecTimeout       time.Duration = time.Second * 2
	defaultExecTimeoutScript time.Duration = time.Second * 5
	runWorkPoolSize                        = 5
	runChannelSize                         = 2 * runWorkPoolSize
)

var (
	cRunChannel    chan RunResponse
	cppRunChannel  chan RunResponse
	javaRunChannel chan RunResponse
	python3Channel chan RunResponse
	python2Channel chan RunResponse
	goRunChannel   chan RunResponse
)

var runChannelMap map[string]chan RunResponse

var runChannels []chan RunResponse

//Maps run commands for languages
var RunCommandMap map[string]func(*Code) *exec.Cmd

//Run command Array
var runCommand []func(*Code) *exec.Cmd

//C exec command: format "PATH+FILENAME"
func c_run(code *Code) *exec.Cmd {
	executable := code.path + code.name
	return exec.Command(executable)
}

//Java exec command: formant "java -cp PATH CLASSNAME"
func java_run(code *Code) *exec.Cmd {
	executor := "java"
	option := "-cp"
	class := "Main"
	return exec.Command(executor, option, code.path, class)
}

//Python 3.4 exec command: format "python3 PATH+FILENAME+.py"
func python3_run(code *Code) *exec.Cmd {
	executor := "python3"
	script := code.path + code.name + ".py"
	return exec.Command(executor, script)
}

//Python2.7 exec command: format "python PATH+FILENAME+.py"
func python2_run(code *Code) *exec.Cmd {
	executor := "python"
	script := code.path + code.name + ".py"
	return exec.Command(executor, script)
}

//C++ exec command : format "PATH+FILENAME"
func cpp_run(code *Code) *exec.Cmd {
	executable := code.path + code.name
	return exec.Command(executable)
}

//Go exec command : format "go run PATH+FILENAME+.py"
func go_run(code *Code) *exec.Cmd {
	toRun := code.path + code.name
	return exec.Command(toRun)
}

//Error checker
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

/*Worker Function
Panics if Unknown language found
Status Codes:
	0 : Execution Success
	1 : Exection timeout
	2 : Undefined error
*/
func run(channel chan RunResponse) {
	for work := range channel {
		code := work.code
		cmd := RunCommandMap[code.Lang](code)
		stdinPipe, _ := cmd.StdinPipe()
		stdoutPipe, _ := cmd.StdoutPipe()
		stderrPipe, _ := cmd.StderrPipe()
		io.WriteString(stdinPipe, code.Stdin)
		localChan := make(chan int)
		var (
			outByte []byte
			errByte []byte
			err     error
		)
		go func() {
			cmd.Start()
			outByte, _ = ioutil.ReadAll(stdoutPipe)
			errByte, _ = ioutil.ReadAll(stderrPipe)
			err = cmd.Wait()
			localChan <- 1
		}()
		select {
		case <-localChan:
			if err != nil {
				if strings.Contains(err.Error(), "segmentation") {
					code.Stderr = "Segmentation Fault"
					break
				}
			}
			if len(errByte) != 0 {
				code.Stderr = string(errByte)
			} else {
				code.Stdout = string(outByte)
				code.RunStatus = true
			}
		case <-time.After(code.execTimeLimit):
			cmd.Process.Kill()
			log.Println("Run Timeout Lang: ", code.Lang)
			code.Stderr = "Execution Timeout"
		}
		work.runResponse <- 0
	}
}

/*
func run(statusChannel chan int, code *Code) {
	defer func() {
		if err := recover(); err != nil {
			log.Print("Run Error: Lang:", code.Lang)
			statusChannel <- statusUnknown
		}
	}()
	function, e := RunCommandMap[code.Lang]
	if e != true {
		panic("Run: Panic Encountered Lang:" + code.Lang)
	}
	cmd := function(code)
	stdinPipe, err := cmd.StdinPipe()
	CheckError(err)
	stdoutPipe, err := cmd.StdoutPipe()
	CheckError(err)
	stderrPipe, err := cmd.StderrPipe()
	CheckError(err)
	cmd.Start()
	go func() {
		io.WriteString(stdinPipe, code.Stdin)
		stdinPipe.Close()
	}()
	go func() {
		bytes, _ := ioutil.ReadAll(stdoutPipe)
		code.Stdout = string(bytes)
		stdoutPipe.Close()
	}()
	go func() {
		bytes, _ := ioutil.ReadAll(stderrPipe)
		code.Stderr = string(bytes)
		stderrPipe.Close()
	}()
	channel := make(chan int)
	go func() {
		cmd.Wait()
		channel <- statusPass
		code.ExecTime = cmd.ProcessState.UserTime().Seconds()
	}()
	// If channel responds, execution success or failure is found by contents of stderr
	// If statusChannel responds, exection abort
	select {
	case <-channel:
		if code.Stderr != "" {
			statusChannel <- statusFail
		} else {
			statusChannel <- statusPass
		}
	case <-statusChannel:
		cmd.Process.Kill()
		statusChannel <- statusFail
	}
}
*/

func (code *Code) RunManager() {
	res := RunResponse{code: code, runResponse: make(chan int)}
	langChannel := runChannelMap[code.Lang]
	langChannel <- res
	<-res.runResponse
}

func (cr *CRManager) RunBatchManager() {
	code := cr.Program
	res := RunResponse{code: code, runResponse: make(chan int)}
	langChannel := runChannelMap[code.Lang]
	code.execTimeLimit = defaultExecTimeout
	status := TestCaseStatus{}
	for stdin := range cr.Stdin {
		code.Stdin = stdin
		langChannel <- res
		<-res.runResponse
		status.ExecTime = code.ExecTime
		if code.RunStatus == true {
			if code.Stdout == <-cr.Stdout {
				status.Success = true
				status.Comment = "Correct"
				cr.Status <- status
			} else {
				status.Success = false
				status.Comment = "Wrong Answer"
				cr.Status <- status
			}
		} else {
			status.Success = false
			status.Comment = code.Stderr
			<-cr.Stdout
			cr.Status <- status
		}
	}
}

/*
//Execution Manager
func (code *Code) RunManager() {
	var execTime time.Duration
	statusChannel := make(chan int)
	if code.execTimeLimit == 0 {
		execTime = defaultExecTimeout
	} else {
		execTime = code.execTimeLimit
	}
	go run(statusChannel, code)
	select {
	case response := <-statusChannel:
		if response == statusPass {
			code.RunStatus = true
		} else if response == statusFail {
			log.Println("Run Error: Run Failed")
		}
	case <-time.After(time.Second * execTime):
		log.Println("Run : Execution Timeout")
		statusChannel <- statusUnknown
		<-statusChannel
	}
}
*/
func init() {

	runChannelMap = make(map[string]chan RunResponse)
	RunCommandMap = make(map[string]func(*Code) *exec.Cmd)
	//	supportedLangs := []string{"C", "C++", "Java", "Python3", "Python2", "Go"}
	runChannels = []chan RunResponse{cRunChannel, cppRunChannel,
		javaRunChannel, python3Channel,
		python2Channel, goRunChannel}
	for iter := range runChannels {
		runChannels[iter] = make(chan RunResponse, runChannelSize)
	}
	runCommand := []func(code *Code) *exec.Cmd{
		c_run, cpp_run,
		java_run, python3_run,
		python2_run, go_run}
	for iter, lang := range supportedLangs {
		RunCommandMap[lang] = runCommand[iter]
		runChannelMap[lang] = runChannels[iter]
	}
	for _, channel := range runChannels {
		for i := 0; i < runWorkPoolSize; i++ {
			go run(channel)
		}
	}
	log.Println("Run Init: Normal")
}
