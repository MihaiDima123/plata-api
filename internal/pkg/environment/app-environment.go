package environment

type AppEnvironment struct {
	DatabaseHost string
	DatabaseUsername string
	DatabasePassword string
	DatabaseName string
}

type ConfigLoader interface {
	Load() AppEnvironment
}