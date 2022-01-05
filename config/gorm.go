package config

import (
	"context"
	"database/sql"
	"time"

	"github.com/habibiiberahim/go-backend/exception"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDatabase(configuration Config) *gorm.DB {
	dsn := "root:root_root@tcp(127.0.0.1:3306)/lolosasn?charset=utf8mb4&parseTime=True&loc=Local"

	sqlDB, err := sql.Open("mysql", dsn)

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	exception.PanicIfNeeded(err)

	database, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	exception.PanicIfNeeded(err)

	return database
}

func NewGormContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
