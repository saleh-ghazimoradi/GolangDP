package singleton

import (
	"database/sql"
	"log/slog"
	"os"
	"sync"
)

// Using Singleton pattern for the config setup
var (
	configInstance *Config
	configOnce     sync.Once
)

type Config struct {
	DatabaseURL string
	ApiKey      string
}

func GetConfig() *Config {
	configOnce.Do(func() {
		configInstance = &Config{
			DatabaseURL: "postgres://postgres:postgres@localhost/postgres?sslmode=disable",
			ApiKey:      "your-api-key",
		}
	})
	return configInstance
}

// Using Singleton pattern for DB pool setup
var (
	dbInstance *DBPool
	dbOnce     sync.Once
)

type DBPool struct {
	Db *sql.DB
}

func GetDBPool() *DBPool {
	dbOnce.Do(func() {
		db, err := sql.Open("postgres", "")
		if err != nil {
			panic(err)
		}
		dbInstance = &DBPool{
			Db: db,
		}
	})
	return dbInstance
}

// Using Singleton pattern to implement logging service in Golang

var (
	logInstance *Logger
	logOnce     sync.Once
)

type Logger struct {
	Log *slog.Logger
}

func GetLogger() *Logger {
	logOnce.Do(func() {
		logInstance = &Logger{Log: slog.New(slog.NewJSONHandler(os.Stdout, nil))}
	})
	return logInstance
}
