package config

import (
	"flag"

	"github.com/zeromicro/go-zero/core/conf"
)

func GetTestConf(cfile string) Config {
	flag.Parse()
	var c Config
	conf.MustLoad(cfile, &c)

	return c
}
