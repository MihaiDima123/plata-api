package database

import "plata-ui/internal/pkg/environment"

type DbConfiguration struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
}

type DbConfigurationLoader interface {
	Load(environment.AppEnvironment) DbConfiguration
}
