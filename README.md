# goject - x86/x64 Reflective Dll Injector

COMING SOON

Reflectively load x86/x64 from memory

Please note this is a VERY basic & bare bones reflective injector.

# Installation

Below is how you install it as a library
```
go get github.com/0xffffa/goject
```

And here's how you install it as a cmd to inject a dll to a process
```
go install github.com/0xffffa/goject/cmd/injector
```

# Usage

Here's an example of using it as a command:
```
goject <dll> <process_name>
```

Lastly, here's an example program:
```go
package main

import (
    loader "github.com/0xffffa/goject"
    "github.com/mitchellh/go-ps"
)

func main() {
    dllData, _ := ioutil.ReadFile("test1.dll")
    if err := loader.Inject(findProcess("csgo.exe"), dllData); err != nil {
        panic(err)
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

```
