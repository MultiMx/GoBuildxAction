package tools

import (
	"fmt"
	"os"
)

func SetOutput(name, value string) {
	file, e := os.OpenFile(os.Getenv("GITHUB_OUTPUT"), os.O_APPEND|os.O_WRONLY, 0600)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	_, e = file.WriteString(fmt.Sprintf("%s=%s", name, value) + "\n")
	if e != nil {
		panic(e)
	}
}
