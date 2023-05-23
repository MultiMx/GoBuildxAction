package goBuild

import (
	"fmt"
	"strings"
)

func ArgsFromString(args string) []string {
	return strings.Split(args, " ")
}

type Arch struct {
	OS   string
	ARCH string
}

func (a Arch) Env() []string {
	return []string{"GOOS=" + a.OS, "GOARCH=" + a.ARCH}
}

func (a Arch) OutputArgs(name string) []string {
	return []string{"-o", fmt.Sprintf("%s/%s/%s", a.OS, a.ARCH, name)}
}

func (a Arch) String() string {
	return a.OS + "/" + a.ARCH
}
