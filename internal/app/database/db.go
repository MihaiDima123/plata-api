package appDatabase

import (
	"fmt"
	"plata-ui/internal/pkg/database"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConn = &gorm.DB{}

func ConfigureDb(config *database.DbConfiguration) error {
	dsn := getDsn(config)
	conn, gormErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if gormErr != nil {
		return gormErr
	}

	DbConn = conn

	sqlDB, sqlErr := conn.DB()
	if sqlErr != nil {
		return sqlErr
	}

	sqlDB.SetMaxIdleConns(config.MaxIdle)
	sqlDB.SetMaxOpenConns(config.MaxPoolSize)
	sqlDB.SetConnMaxLifetime(time.Minute)

	return nil
}

func getDsn(config *database.DbConfiguration) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		config.Host,
		config.Username,
		config.Password,
		config.DatabaseName,
		config.DatabasePort,
	)
}
