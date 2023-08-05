package config

import "go.uber.org/fx"

type Config struct {
	AppName string
	Port    int
}

func NewConfig() *Config {
	return &Config{
		AppName: "",
		Port:    8080,
	}
}

var Module = fx.Module(
	"config",
	fx.Provide(NewConfig),
)
