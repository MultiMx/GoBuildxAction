package main

import (
	"github.com/MultiMx/GoBuildxAction/internal/global"
	"github.com/MultiMx/GoBuildxAction/internal/pkg/goBuild"
	"strings"
)

func main() {
	arches := strings.Split(global.Config.Platform, ",")
	builder := goBuild.NewMultiPlatformBuilder(
		global.Config.Name,
		global.Config.Args,
		global.Config.Target, arches,
	)
	builder.Build()
}
