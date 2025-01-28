package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	loader "github.com/0xffffa/goloader"
	"github.com/mitchellh/go-ps"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: goloader <dll> <process_name>")
		return
	}

	dllPath := os.Args[1]
	processName := os.Args[2]

	if len(dllPath) == 0 || len(processName) == 0 {
		fmt.Println("Usage: goloader <dll> <process_name>")
		return
	}

	dllData, _ := ioutil.ReadFile(dllPath)
	if len(dllData) == 0 {
		fmt.Println("Unable to find dll", dllPath)
		return
	}

	processId := findProcess(processName)

	if processId == 0 {
		fmt.Println("Unable to find process id for process", processName)
		return
	}

	if err := loader.Inject(processId, dllData); err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully injected dll to process", processName)
	}

}

func findProcess(name string) int {
	var foundPid int
	processes, err := ps.Processes()
	if err != nil {
		return foundPid
	}
	for _, process := range processes {
		if !strings.Contains(strings.ToLower(process.Executable()), name) {
			continue
		}
		foundPid = process.Pid()
	}
	return foundPid
}
