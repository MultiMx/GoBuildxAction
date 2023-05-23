package global

import (
	"github.com/MultiMx/GoBuildxAction/internal/global/models"
	"github.com/caarlos0/env/v6"
	"log"
)

var Config models.Config

func init() {
	if e := env.Parse(&Config, env.Options{
		Prefix: "INPUT_",
	}); e != nil {
		log.Fatalln(e)
	}
}
