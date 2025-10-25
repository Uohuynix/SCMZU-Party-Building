package db

import (
	"database/sql"
	"time"

	"SCMZU_Party_Building/config"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB returns (*gorm.DB, *sql.DB, error)
func InitDB(conf *config.Config) (*gorm.DB, *sql.DB, error) {
	// First create native sql.DB (with multiStatements enabled) for running SQL scripts
	sqlDSN := conf.MySQLDSNWithMultiStatements()
	sqlDB, err := sql.Open("mysql", sqlDSN)
	if err != nil {
		return nil, nil, err
	}
	// Configure connection pool
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// ping to verify connectivity
	if err := sqlDB.Ping(); err != nil {
		return nil, nil, err
	}

	// Then create GORM DB (separate connection).
	gormDSN := conf.MySQLDSN()
	gormDB, err := gorm.Open(mysql.Open(gormDSN), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		_ = sqlDB.Close()
		return nil, nil, err
	}

	return gormDB, sqlDB, nil
}
