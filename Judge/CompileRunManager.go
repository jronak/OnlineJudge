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
	crOnceWorkSize      = 20
	crOnceChanBufSize   = 2 * crOnceWorkSize
)

type CRManager struct {
	Program       *Code
	RawCode       *string
	Pid           int
	IsCustomInput bool
	IsBatch       bool
	batchSize     int
	stdin         chan string
	stdout        chan string
	status        chan bool
	receive       chan int
}

var crChannelOnce chan *CRManager

//All the Languages supported listed here
var supportedLangs []string

/*
	Buffered Channel holds the path names for Java programs : Buffer Size :10
	Java Code should contain Main.class.
	Main.class can be disturbed by other routines
	Hence separate DIR for every java programa
*/
var javaPathChan chan string

/*
	Buffered Channel holds the file names : Buffer Size : 1000
	Maintains unique filenames
*/
var fileNameChan chan string

//	Code Extensions : {".c",".java",...}
var codeExtensionsMap map[string]string

//Bool: if the program is compilable true
func (cr *CRManager) isCompilableLang() bool {
	_, val := CompileCommandMap[cr.Program.Lang]
	return val
}

//Dumps the program into a file
func (cr *CRManager) createCode() {
	path := cr.Program.path + cr.Program.name + codeExtensionsMap[cr.Program.Lang]
	err := ioutil.WriteFile(path, []byte(*cr.RawCode), fileDirPermissions)
	CheckError(err)
}

//Deletes program file
func (cr *CRManager) deleteCode() {
	if cr.isCompilableLang() {
		os.Remove(cr.Program.path + cr.Program.name + codeExtensionsMap[cr.Program.Lang])
	}
}

//Deletes the Executables
func (cr *CRManager) deleteExec() {
	if cr.Program.Lang == "Java" {
		os.RemoveAll(cr.Program.path)
	} else if cr.isCompilableLang() {
		os.Remove(cr.Program.path + cr.Program.name)
	} else {
		os.Remove(cr.Program.path + cr.Program.name + codeExtensionsMap[cr.Program.Lang])
	}
}

//Sets the path for Program
func (cr *CRManager) setPath() {
	if cr.Program.Lang == "Java" {
		cr.Program.path = <-javaPathChan
	} else {
		cr.Program.path = defaultPath
	}
}

//Adds the file name back to the buffer
func (cr *CRManager) resetName() {
	fileNameChan <- cr.Program.name
}

//Sets the name
func (cr *CRManager) setName() {
	cr.Program.name = <-fileNameChan
}

//Adds the path back to the buffer
func (cr *CRManager) resetPath() {
	if cr.Program.Lang == "Java" {
		generatePaths(1)
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
func crOnceWorker(channel chan *CRManager) {
	for cr := range channel {
		run := true
		cr.setName()
		cr.setPath()
		cr.createCode()
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
		func() {
			cr.deleteCode()
			cr.deleteExec()
			cr.resetPath()
			cr.resetName()
		}()
	}
}

func (cr *CRManager) CROnce() {
	cr.receive = make(chan int)
	crChannelOnce <- cr
	<-cr.receive
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

//** To be defered in the Main
func removeTempPath() {
	os.RemoveAll(rootPath)
}

func init() {
	javaPathChan = make(chan string, javaPathsBufferSize)
	fileNameChan = make(chan string, nameBufferSize)
	crChannelOnce = make(chan *CRManager, crOnceChanBufSize)
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
	}
	log.Println("CRManager Init : Normal")
}
