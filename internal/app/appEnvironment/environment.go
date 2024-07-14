package appEnvironment

import (
	"os"
	"plata-ui/internal/pkg/environment"
	"strconv"

	"github.com/joho/godotenv"
)

var Env = &environment.AppEnvironment{}

func Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	// Database
	Env.DatabaseHost = os.Getenv("DATABASE_HOST")
	Env.DatabaseName = os.Getenv("DATABASE_NAME")
	Env.DatabasePort, _ = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	Env.DatabaseUsername = os.Getenv("DATABASE_USERNAME")
	Env.DatabasePassword = os.Getenv("DATABASE_PASSWORD")
	Env.DatabaseMaxIdleConn, _ = strconv.Atoi(os.Getenv("DATABASE_POOL_MAX_IDLE_CONN"))
	Env.DatabaseMaxOpenConn, _ = strconv.Atoi(os.Getenv("DATABASE_POOL_MAX_OPEN_CONN"))

	return nil
}
