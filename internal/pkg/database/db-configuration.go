package database

import (
	"plata-ui/internal/pkg/environment"
)

type DbConfiguration struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	DatabasePort int

	MaxIdle     int
	MaxPoolSize int
}

func LoadFromApp(env *environment.AppEnvironment) *DbConfiguration {
	return &DbConfiguration{
		Host:         env.DatabaseHost,
		Username:     env.DatabaseUsername,
		Password:     env.DatabasePassword,
		DatabaseName: env.DatabaseName,
		DatabasePort: env.DatabasePort,
		MaxIdle:      env.DatabaseMaxIdleConn,
		MaxPoolSize:  env.DatabaseMaxOpenConn,
	}
}
