package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	JetAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	Database struct {
		Postgres struct {
			Dsn         string
			IsMigration bool
		}
	}
}
