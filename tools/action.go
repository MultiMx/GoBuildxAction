package tools

import (
	"fmt"
	"os"
	"strings"
)

func SetOutput(name, value string) {
	value = strings.ReplaceAll(value, "\"", `\"`)
	file, e := os.OpenFile(os.Getenv("GITHUB_OUTPUT"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if e != nil {
		panic(e)
	}
	defer file.Close()

	_, e = file.WriteString(fmt.Sprintf("%s=%s\n", name, value))
	if e != nil {
		panic(e)
	}
}
