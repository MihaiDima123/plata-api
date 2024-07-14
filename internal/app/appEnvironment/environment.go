package appEnvironment

import (
	"fmt"
	"os"
	"plata-ui/internal/pkg/environment"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const (
	ERRORS_EMPTY_TEMPLATE = "Empty String for %s"
)

var Env = &environment.AppEnvironment{}

func Load() error {
	var err error = nil

	if err := godotenv.Load(); err != nil {
		return err
	}

	// Database
	Env.DatabaseHost, err = getString("DATABASE_HOST")
	Env.DatabaseName, err = getString("DATABASE_NAME")
	Env.DatabasePort, err = getNumber("DATABASE_PORT")
	Env.DatabaseUsername, err = getString("DATABASE_USERNAME")
	Env.DatabasePassword, err = getString("DATABASE_PASSWORD")
	Env.DatabaseMaxIdleConn, err = getNumber("DATABASE_POOL_MAX_IDLE_CONN")
	Env.DatabaseMaxOpenConn, err = getNumber("DATABASE_POOL_MAX_OPEN_CONN")

	return err
}

func getString(name string) (string, error) {
	value := os.Getenv(name)

	if strings.Trim(value, "") == "" {
		return value, fmt.Errorf(ERRORS_EMPTY_TEMPLATE, name)
	}

	return value, nil
}

func getNumber(name string) (int, error) {
	return strconv.Atoi(os.Getenv(name))
}
