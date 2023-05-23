package goBuild

import (
	"log"
	"strings"
)

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
