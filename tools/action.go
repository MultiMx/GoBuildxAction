package tools

import (
	"fmt"
	"os/exec"
	"strings"
)

func SetOutput(name, value string) {
	value = strings.ReplaceAll(value, "\"", `\"`)
	e := exec.Command("/bin/sh", "-c", fmt.Sprintf("echo \"%s=%s\" >> $GITHUB_OUTPUT", name, value)).Run()
	if e != nil {
		panic(e)
	}
}
