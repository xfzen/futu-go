package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Database struct {
		Driver string // 驱动名，默认为sqlite，支持mysql
		Source string // Datasource
	}
}
