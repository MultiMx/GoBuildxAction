package goBuild

import (
	"fmt"
	"log"
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

func NewMultiPlatformBuilder(name, args, target string, arches []string) MultiPlatformBuilder {
	var targetArches = make([]Arch, len(arches))
	for i, arch := range arches {
		t := strings.Split(arch, "/")
		if len(t) != 2 {
			log.Fatalln("illegal platform", arch)
		}
		targetArches[i].OS, targetArches[i].ARCH = t[0], t[1]
	}
	return MultiPlatformBuilder{
		arches: targetArches,
		args:   ArgsFromString(args),
		name:   name,
		target: target,
	}
}

type MultiPlatformBuilder struct {
	arches []Arch
	args   []string
	name   string
	target string
}

func (a MultiPlatformBuilder) Commands() []string {
	var commands = make([]string, len(a.arches))
	for i, arch := range a.arches {
		args := append(arch.Env(), append([]string{"go build"}, append(a.args, append(arch.OutputArgs(a.name), a.target)...)...)...)
		commands[i] = strings.Join(args, " ")
	}
	return commands
}
