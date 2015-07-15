package Judge

import (
	"log"
	"os/exec"
	"time"
)

/* 	Struct used throughout compilation
path : Absolute path of the parent of the file 		Format : "/home/user/Desktop/"
name : name of the file to be compiled		Format : "Main"
Lang : Defines language
Stderr :  Error
Stdin : Input
Stdout : Output
CompilationStatus : True if compilation was successful else false
RunStatus : True if run was Successful else false
To-Do
**Addons:
	*
*/
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

type CompilerResponse struct {
	code             *Code
	compilerResponse chan int
}

var (
	cChannel    chan CompilerResponse
	javaChannel chan CompilerResponse
	cppChannel  chan CompilerResponse
)

const (
	statusPass        = 0
	statusFail        = 1
	statusUnknown     = 2
	compilationTimout = 10 //2 seconds
	workerPoolSize    = 10
	channelBufferSize = 2 * workerPoolSize
)

var compileChannelMap map[string]chan CompilerResponse

var compileChannels []chan CompilerResponse

//	Map of different compilation command functions
var CompileCommandMap map[string]func(code *Code) *exec.Cmd

//	Array of Compilation command functions
var compileCommands []func(code *Code) *exec.Cmd

/***Slice of languages supported **
Caution: Languages added here should have compilation command
**/
var compilableLangs []string

/*
C Compilation statement
GCC Version : 4.8.2
External libs :  Math(-lm)
*/
func c_command(code *Code) *exec.Cmd {
	compiler := "gcc"
	options := "-o"
	extLibs := "-lm"
	objPath := code.path + code.name
	codePath := code.path + code.name + ".c"
	return exec.Command(compiler, options, objPath, codePath, extLibs)
}

/*
Java Complation statement
Javac Version: Open-jdk 7
External libs : -
*/
func java_command(code *Code) *exec.Cmd {
	compiler := "javac"
	//	Destination for the class files
	options := "-d"
	objPath := code.path
	codePath := code.path + code.name + ".java"
	return exec.Command(compiler, options, objPath, codePath)
}

/*
C Compilation statement
GCC Version : 4.8.2
External libs :
*/
func cpp_command(code *Code) *exec.Cmd {
	compiler := "g++"
	options := "-o"
	objPath := code.path + code.name
	codePath := code.path + code.name + ".cpp"
	return exec.Command(compiler, options, objPath, codePath)
}

// 	Compiles the code
// 	Handled by CompilationManager()
/*
func compile(statusChannel chan int, code *Code) {

	defer func() {
		if err := recover(); err != nil {
			code.Stdout = "Oops!, Internal Error"
			statusChannel <- statusUnknown
		}
	}()
	function, e := CompileCommandMap[code.Lang]
	if e != true {
		panic("Compilation Panic Lang: " + code.Lang)
	}
	cmd := function(code)
	out, err := cmd.CombinedOutput()
	if err != nil {
		code.Stderr = string(out)
		log.Println("Compiler Error: ", err)
		statusChannel <- statusFail
	} else {
		code.Stdout = string(out)
		statusChannel <- statusPass
	}

}
*/

func compile(channel chan CompilerResponse) {
	for work := range channel {
		code := work.code
		localChan := make(chan int)
		cmd := CompileCommandMap[code.Lang](code)
		var out []byte
		var err error
		go func() {
			out, err = cmd.CombinedOutput()
			localChan <- 0
		}()
		select {
		case <-localChan:
			if err == nil {
				code.CompilationStatus = true
			}
			code.Stderr = string(out)
		case <-time.After(time.Second * compilationTimout):
			code.Stderr = "Compilation Timeout"
			log.Println("Compilation Timeout Lang: ", code.Lang)
		}
		work.compilerResponse <- 0
	}
}

/*	Code method
	Handles compilation of the supported languages
	Maintains a statusChannel (chan int) between compile function

	Compilation Timeout := 5 Seconds

	Compilation CompilationStatus code:
	0 - Success
	1 - Error in compilation
	2 - Internal Error
*/
/*
func (code *Code) CompilationManager() {
	statusChannel := make(chan int)
	go compile(statusChannel, code)
	select {
	case val := <-statusChannel:
		if val == statusPass {
			code.CompilationStatus = true
		} else if val == statusUnknown {
			log.Println("Compiler Error: Unknow Language:" + code.Lang)
		}
	case <-time.After(time.Second * compilationTimout):
		code.Stderr = "Oops, compilation timeout"
		log.Println("Compiler Error: Compilation Timeout")
	}
}
*/

func (code *Code) CompilationManager() {
	res := CompilerResponse{code: code, compilerResponse: make(chan int)}
	channel := compileChannelMap[code.Lang]
	channel <- res
	<-res.compilerResponse

}

func init() {

	javaChannel = make(chan CompilerResponse, channelBufferSize)
	cChannel = make(chan CompilerResponse, channelBufferSize)
	cppChannel = make(chan CompilerResponse, channelBufferSize)
	compileChannels = [](chan CompilerResponse){cChannel, javaChannel, cppChannel}
	compileChannelMap = make(map[string]chan CompilerResponse)
	CompileCommandMap = make(map[string]func(*Code) *exec.Cmd)
	compilableLangs = []string{"C", "Java", "C++"}
	compileCommands = [](func(code *Code) *exec.Cmd){
		c_command,
		java_command,
		cpp_command}
	for iter, Lang := range compilableLangs {
		CompileCommandMap[Lang] = compileCommands[iter]
		compileChannelMap[Lang] = compileChannels[iter]
	}

	for _, channel := range compileChannels {
		for i := 0; i < workerPoolSize; i++ {
			go compile(channel)
		}
	}

	log.Println("Compiler Init: Normal")
}
