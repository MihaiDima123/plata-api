package main

import (
	"fmt"
	"plata-ui/internal/app/appEnvironment"
	appDatabase "plata-ui/internal/app/database"
	"plata-ui/internal/pkg/database"
)

func main() {
	// -- App env
	if err := appEnvironment.Load(); err != nil {
		panic(fmt.Sprintf("Could not load app environment, %s", err.Error()))
	}

	// -- Database
	databaseConfig := database.LoadFromApp(appEnvironment.Env)
	if err := appDatabase.ConfigureDb(databaseConfig); err != nil {
		panic(fmt.Sprintf("Could not configure the dB, %s", err.Error()))
	}

}
