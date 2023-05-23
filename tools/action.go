package tools

import (
	"fmt"
	"os"
	"os/exec"
)

func SetOutput(name, value string) {
	_ = exec.Command("echo", fmt.Sprintf("%s=%s", name, value), ">>", os.Getenv("GITHUB_OUTPUT")).Run()
}
