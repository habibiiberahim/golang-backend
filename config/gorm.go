package config

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/habibiiberahim/go-backend/entity"
	"github.com/habibiiberahim/go-backend/exception"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDatabase(configuration Config) *gorm.DB {

	user := configuration.Get("DB_USER")
	pass := configuration.Get("DB_PASS")
	host := configuration.Get("DB_HOST")
	port := configuration.Get("DB_PORT")
	dbname := configuration.Get("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)

	sqlDB, err := sql.Open("mysql", dsn)
	exception.PanicIfNeeded(err)

	maxIdleConn, err := strconv.Atoi(configuration.Get("DB_MAX_IDLE_CONNECTION"))
	exception.PanicIfNeeded(err)

	maxOpenConn, err := strconv.Atoi(configuration.Get("DB_MAX_OPEN_CONNECTION"))
	exception.PanicIfNeeded(err)

	maxLifetimeConn, err := strconv.Atoi(configuration.Get("DB_MAX_LIFETIME_CONNECTION"))
	exception.PanicIfNeeded(err)

	sqlDB.SetMaxIdleConns(maxIdleConn)

	sqlDB.SetMaxOpenConns(maxOpenConn)

	sqlDB.SetConnMaxLifetime(time.Duration(maxLifetimeConn) * time.Minute)

	database, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	exception.PanicIfNeeded(err)

	database.AutoMigrate(&entity.User{})

	return database
}

func NewGormContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
