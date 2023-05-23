package models

type Config struct {
	Name string `env:"NAME,required"`
	Args string `env:"ARGS"`
}
