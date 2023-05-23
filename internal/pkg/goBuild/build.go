package goBuild

import (
	"os"
	"os/exec"
	"strings"
)

func ArgsFromString(args string) []string {
	return strings.Split(args, " ")
}

func Execute(env []string, args ...string) error {
	args = append([]string{"build"}, args...)
	cmd := exec.Command("go", args...)
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
