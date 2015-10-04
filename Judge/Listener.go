package Judge

import (
	"encoding/json"
	"log"
	"net/http"
	//	"os"
	"strconv"
)

const (
	directories     = 5
	prefix          = "d"
	requestName     = "problem.json"
	solutionName    = "solution.json"
	ProgramFileName = "code"
)

var (
	dirString chan string
)

func worker(dir string, js string) string {
	defer func() {
		if x := recover(); x != nil {
			log.Println("Recovered from shit")
		}
	}()
	cr := CRManager{}
	json.Unmarshal([]byte(js), &cr)
	cr.Program.name = ProgramFileName
	cr.Program.path = dir
	cr.CR()
	solutionBytes, _ := json.Marshal(&cr)
	dirString <- dir
	return string(solutionBytes)
}

func serve(writer http.ResponseWriter, response *http.Request) {
	response.ParseForm()
	prob := response.Form.Get("json")
	solu := worker(<-dirString, prob)
	writer.Header().Add("json", solu)
}

func Start() {
	dirString = make(chan string, 10)
	log.Println("Judge: Successful")
	for i := 1; i <= directories; i++ {
		dir := prefix + strconv.Itoa(i) + "/"
		dirString <- dir
	}
	http.HandleFunc("/", serve)
	http.ListenAndServe(":8080", nil)
}
