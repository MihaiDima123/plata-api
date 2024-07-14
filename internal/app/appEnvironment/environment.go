package appEnvironment

import (
	"errors"
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
	if err := godotenv.Load(); err != nil {
		return err
	}

	// Database
	Env.DatabaseHost = os.Getenv("DATABASE_HOST")
	Env.DatabaseName = os.Getenv("DATABASE_NAME")

	dbPort, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	Env.DatabasePort = dbPort
	if err != nil {
		return err
	}

	Env.DatabaseUsername = os.Getenv("DATABASE_USERNAME")
	Env.DatabasePassword = os.Getenv("DATABASE_PASSWORD")

	dbMaxIdle, err := strconv.Atoi(os.Getenv("DATABASE_POOL_MAX_IDLE_CONN"))
	Env.DatabaseMaxIdleConn = dbMaxIdle
	if err != nil {
		return err
	}

	dbPoolSize, err := strconv.Atoi(os.Getenv("DATABASE_POOL_MAX_OPEN_CONN"))
	Env.DatabaseMaxOpenConn = dbPoolSize
	if err != nil {
		return err
	}

	return nil
}

func getString(val string, name string) (string, error) {
	value := os.Getenv(name)

	if strings.Trim(val, "") == "" {
		return val, errors.New(fmt.Sprintf(ERRORS_EMPTY_TEMPLATE, name))
	}

	return value, nil
}

func getNumber() (int, error) {
	return strconv.Atoi(os.Getenv("DATABASE_HOST"))
}
