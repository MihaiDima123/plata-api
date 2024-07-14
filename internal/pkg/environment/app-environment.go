package environment

type AppEnvironment struct {
	DatabaseHost     string
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
	DatabasePort     int
	DatabaseMaxIdleConn int
	DatabaseMaxOpenConn int
}
