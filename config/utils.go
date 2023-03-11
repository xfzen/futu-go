package config

import (
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("zf", "etc/futuq-api.yaml", "the config file")

func LoadConf(path string) Config {
	var c Config
	conf.MustLoad(path, &c)
	return c
}

func GetConf() Config {
	flag.Parse()
	var c Config
	conf.MustLoad(*configFile, &c)

	return c
}
