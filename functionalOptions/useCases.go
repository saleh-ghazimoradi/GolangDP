package functionalOptions

import (
	"fmt"
	"io"
	"os"
	"time"
)

// DB configuration

type DBClient struct {
	DBHost          string
	DBPort          int
	DBMaxPoolSize   int
	DBTimeout       time.Duration
	DBRetryAttempts int
}

type OptionDb func(*DBClient)

func WithDBHost(host string) OptionDb {
	return func(db *DBClient) {
		db.DBHost = host
	}
}

func WithDBPort(port int) OptionDb {
	return func(db *DBClient) {
		db.DBPort = port
	}
}

func WithDBMaxPoolSize(maxPoolSize int) OptionDb {
	return func(db *DBClient) {
		db.DBMaxPoolSize = maxPoolSize
	}
}

func WithDBTimeout(timeout time.Duration) OptionDb {
	return func(db *DBClient) {
		db.DBTimeout = timeout
	}
}

func WithDBRetryAttempts(attempts int) OptionDb {
	return func(db *DBClient) {
		db.DBRetryAttempts = attempts
	}
}

func (c *DBClient) Connect() {
	fmt.Printf("Connecting to database at %s:%d\n", c.DBHost, c.DBPort)
	fmt.Printf("Max pool size: %d, Timeout: %v, Retry attempts: %d\n", c.DBMaxPoolSize, c.DBTimeout, c.DBRetryAttempts)
}

func NewDBClient(options ...OptionDb) *DBClient {
	client := &DBClient{
		DBHost:          "localhost",
		DBPort:          5432,
		DBMaxPoolSize:   10,
		DBTimeout:       5 * time.Second,
		DBRetryAttempts: 3,
	}
	for _, opts := range options {
		opts(client)
	}
	return client
}

// Logger implementation using functional options pattern

type Logger struct {
	Level  string
	Output io.Writer
	Prefix string
}

type OptionLogger func(logger *Logger)

func WithLevel(level string) OptionLogger {
	return func(logger *Logger) {
		logger.Level = level
	}
}

func WithOutput(output io.Writer) OptionLogger {
	return func(logger *Logger) {
		logger.Output = output
	}
}

func WithPrefix(prefix string) OptionLogger {
	return func(logger *Logger) {
		logger.Prefix = prefix
	}
}

func (l *Logger) Log(message string) {
	fmt.Fprintf(l.Output, "[%s] %s%s\n", l.Level, l.Prefix, message)
}

func NewLogger(options ...OptionLogger) *Logger {
	logger := &Logger{
		Level:  "INFO",
		Output: os.Stdout,
		Prefix: "",
	}

	for _, opts := range options {
		opts(logger)
	}

	return logger
}
