package Judge

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	//"os/exec"
)

const (
	defaultPath         = "/tmp/OnlineJ/Default/"
	javaPaths           = "/tmp/OnlineJ/JavaP/"
	rootPath            = "/tmp/OnlineJ"
	fileDirPermissions  = 0755
	javaPathsBufferSize = 100
	nameBufferSize      = 1000
	dirSuffixLen        = 2
	fileNameLen         = 6
	crOnceWorkSize      = 2
	crOnceChanBufSize   = 2 * crOnceWorkSize
	deleteBufferSize    = 2 * crOnceChanBufSize
)

type CRManager struct {
	Program *Code
	RawCode *string
	Stdin   chan string
	Stdout  chan string
	Status  chan TestCaseStatus
	Receive chan int
}

/*
	crChannelOnce : Buffered channel for single run requests
	supportedLangs : All the Languages supported listed here
	javaPathChan : Buffered Channel holds the path names for Java programs
	fileNameChan : Buffered Channel holds the file names
	codeExtensionMap :Code Extensions : {".c",".java",...}
*/
var (
	crChannelOnce     chan *CRManager
	crChannelBatch    chan *CRManager
	supportedLangs    []string
	javaPathChan      chan string
	fileNameChan      chan string
	codeExtensionsMap map[string]string
	fileDeleteChan    chan []string
	dirDeleteChan     chan string
)

//Bool: if the program is compilable true
func (cr *CRManager) isCompilableLang() bool {
	_, val := CompileCommandMap[cr.Program.Lang]
	return val
}

//Dumps the program into a file
func (cr *CRManager) createFile() {
	path := cr.Program.path + cr.Program.name + codeExtensionsMap[cr.Program.Lang]
	err := ioutil.WriteFile(path, []byte(*cr.RawCode), fileDirPermissions)
	CheckError(err)
}

//Sets the path for Program
func (cr *CRManager) setPath() {
	if cr.Program.Lang == "Java" {
		cr.Program.path = <-javaPathChan
	} else {
		cr.Program.path = defaultPath
	}
}

//Sets the name
func (cr *CRManager) setName() {
	cr.Program.name = <-fileNameChan
}

func (cr *CRManager) garbageCollector() {
	if cr.Program.Lang == "java" {
		dirDeleteChan <- cr.Program.path
	} else {
		names := []string{cr.Program.path, cr.Program.name, codeExtensionsMap[cr.Program.Lang]}
		fileDeleteChan <- names
	}
}

/*
	Handles:
		Assigns and reuses Names and paths(Java)
		Writes data onto a file
		Cleaning:
			Deletes Code file
			Deletes Executables(if not batch run)
*/

// Worker for Compile once
func crOnceWorker(channel chan *CRManager) {
	for cr := range channel {
		run := true
		cr.setName()
		cr.setPath()
		cr.createFile()
		if cr.isCompilableLang() {
			cr.Program.CompilationManager()
			if cr.Program.CompilationStatus == false {
				cr.Receive <- 1
				run = false
			}
		}
		if run {
			cr.Program.RunManager()
			cr.Receive <- 1
		}
		cr.garbageCollector()
	}
}

// Compiler and run once
func (cr *CRManager) CROnce() {
	cr.Receive = make(chan int)
	if cr.Program.execTimeLimit == 0 {
		if cr.isCompilableLang() {
			cr.Program.execTimeLimit = defaultExecTimeout
		} else {
			cr.Program.execTimeLimit = defaultExecTimeoutScript
		}
	}
	crChannelOnce <- cr
	<-cr.Receive
}

func crBatchWorker(channel chan *CRManager) {
	for cr := range channel {
		run := true
		cr.setName()
		cr.setPath()
		cr.createFile()
		if cr.isCompilableLang() {
			cr.Program.CompilationManager()
			if cr.Program.CompilationStatus == false {
				cr.Receive <- 1
				run = false
			} else {
				cr.Receive <- 0
			}
		} else {
			cr.Receive <- 0
		}
		if run {
			cr.RunBatchManager()
		}
		cr.garbageCollector()
	}
}

func (cr *CRManager) CRBatch() {
	cr.Receive = make(chan int, 1)
	cr.Status = make(chan TestCaseStatus, 1)
	cr.Stdin = make(chan string, 1)
	cr.Stdout = make(chan string, 1)
	crChannelBatch <- cr
}

//Generates random strings
func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

//Generates n number of names and sends it to the channel
func generateNames(n int) {

	for i := 0; i < n; i++ {
		fileNameChan <- randSeq(fileNameLen)
	}
}

//Generates n number of paths, send it to the channel and creates directory as well
func generatePaths(n int) {
	for i := 0; i < n; i++ {
		path := javaPaths + randSeq(dirSuffixLen) + "/"
		os.MkdirAll(path, fileDirPermissions)
		javaPathChan <- path
	}
}

// Deletes the program files as well buffers the name channel
func fileDeleteWorker(channel chan []string) {
	for file := range channel {
		os.Remove(file[0] + file[1])
		os.Remove(file[0] + file[1] + file[2])
		fileNameChan <- file[1]
	}
}

// Deletes the java program dir and generates new path
func dirDeleteWorker(channel chan string) {
	for dir := range channel {
		os.RemoveAll(dir)
		generatePaths(1)
	}
}

//** To be defered in the Main
func removeTempPath() {
	os.RemoveAll(rootPath)
}

func init() {
	javaPathChan = make(chan string, javaPathsBufferSize)
	fileNameChan = make(chan string, nameBufferSize)
	dirDeleteChan = make(chan string, deleteBufferSize)
	fileDeleteChan = make(chan []string, deleteBufferSize)
	crChannelOnce = make(chan *CRManager, crOnceChanBufSize)
	crChannelBatch = make(chan *CRManager, crOnceChanBufSize)
	os.MkdirAll(defaultPath, fileDirPermissions)
	generatePaths(javaPathsBufferSize)
	generateNames(nameBufferSize)
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
	go dirDeleteWorker(dirDeleteChan)
	go fileDeleteWorker(fileDeleteChan)
	go log.Println("CRManager Init : Normal")
}
