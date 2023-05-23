package main

import (
	"github.com/MultiMx/GoBuildxAction/internal/global"
	"github.com/MultiMx/GoBuildxAction/internal/pkg/goBuild"
	"github.com/MultiMx/GoBuildxAction/tools"
	"strings"
)

func main() {
	arches := strings.Split(global.Config.Platform, ",")
	builder := goBuild.NewMultiPlatformBuilder(
		global.Config.Name,
		global.Config.Args,
		global.Config.Target, arches,
	)
	tools.SetOutput("commands", strings.Join(builder.Commands(), " && "))
}
