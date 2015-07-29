package Judge

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	directories     = 1
	prefix          = "d"
	requestName     = "problem.json"
	solutionName    = "solution.json"
	ProgramFileName = "code"
)

func DirectoryListener(dir string) {
	i := true
	_, _ = os.Create(dir + "file")
	for i == true {
		if file, _ := os.OpenFile(dir+requestName, os.O_RDWR, 0777); file != nil {
			file.Close()
			bytes, _ := ioutil.ReadFile(dir + requestName)
			cr := CRManager{}
			json.Unmarshal(bytes, &cr)
			cr.Program.name = ProgramFileName
			cr.Program.path = dir
			cr.CR()
			solutionBytes, _ := json.Marshal(&cr)
			ioutil.WriteFile(dir+solutionName, solutionBytes, 0755)
			os.Remove(dir + requestName)
		}
		time.Sleep(time.Second)
	}
}

func init() {
	log.Println("Judge: Successful")
	for i := 1; i <= directories; i++ {
		dir := prefix + strconv.Itoa(i) + "/"
		go DirectoryListener(dir)
	}
}
