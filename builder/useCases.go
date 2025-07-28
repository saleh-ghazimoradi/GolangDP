package builder

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// DBConfig has been implemented using builder pattern
type DBConfig struct {
	Host    string
	Port    int
	User    string
	Pass    string
	DbName  string
	SslMode string
	MaxConn int
}

type DBConfigBuilder struct {
	config *DBConfig
}

func (b *DBConfigBuilder) Host(host string) *DBConfigBuilder {
	b.config.Host = host
	return b
}

func (b *DBConfigBuilder) Port(port int) *DBConfigBuilder {
	b.config.Port = port
	return b
}

func (b *DBConfigBuilder) User(user string) *DBConfigBuilder {
	b.config.User = user
	return b
}

func (b *DBConfigBuilder) Pass(pass string) *DBConfigBuilder {
	b.config.Pass = pass
	return b
}
func (b *DBConfigBuilder) DbName(dbName string) *DBConfigBuilder {
	b.config.DbName = dbName
	return b
}
func (b *DBConfigBuilder) SslMode(sslMode string) *DBConfigBuilder {
	b.config.SslMode = sslMode
	return b
}
func (b *DBConfigBuilder) MaxConn(maxConn int) *DBConfigBuilder {
	b.config.MaxConn = maxConn
	return b
}

func (b *DBConfigBuilder) DBBuild() *DBConfig {
	return b.config
}

func NewDBConfigBuilder() *DBConfigBuilder {
	return &DBConfigBuilder{
		config: &DBConfig{
			Host:    "5432",
			SslMode: "disable",
			MaxConn: 10,
		},
	}
}

// SQLQuery has been implemented using builder pattern
type SQLQuery struct {
	table   string
	columns []string
	where   string
	orderBy string
	limit   int
	offset  int
}

type SQLQueryBuilder struct {
	query *SQLQuery
}

func (b *SQLQueryBuilder) Select(columns ...string) *SQLQueryBuilder {
	b.query.columns = columns
	return b
}

func (b *SQLQueryBuilder) From(table string) *SQLQueryBuilder {
	b.query.table = table
	return b
}

func (b *SQLQueryBuilder) Where(where string) *SQLQueryBuilder {
	b.query.where = where
	return b
}

func (b *SQLQueryBuilder) OrderBy(orderBy string) *SQLQueryBuilder {
	b.query.orderBy = orderBy
	return b
}

func (b *SQLQueryBuilder) Limit(limit int) *SQLQueryBuilder {
	b.query.limit = limit
	return b
}

func (b *SQLQueryBuilder) Offset(offset int) *SQLQueryBuilder {
	b.query.offset = offset
	return b
}

func (b *SQLQueryBuilder) Build() string {
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(b.query.columns, ", "), b.query.table)
	if b.query.where != "" {
		query += fmt.Sprintf(" WHERE %s", b.query.where)
	}
	if b.query.orderBy != "" {
		query += fmt.Sprintf(" ORDER BY %s", b.query.orderBy)
	}
	if b.query.limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", b.query.limit)
	}
	if b.query.offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", b.query.offset)
	}
	return query
}

func NewSQLQueryBuilder() *SQLQueryBuilder {
	return &SQLQueryBuilder{
		query: &SQLQuery{},
	}
}

// Server has been implemented using builder pattern
type Server struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Middleware   []string
	TLS          bool
}

type ServerBuilder struct {
	server *Server
}

func (b *ServerBuilder) Port(port int) *ServerBuilder {
	b.server.Port = port
	return b
}

func (b *ServerBuilder) ReadTimeout(readTimeout time.Duration) *ServerBuilder {
	b.server.ReadTimeout = readTimeout
	return b
}

func (b *ServerBuilder) WriteTimeout(writeTimeout time.Duration) *ServerBuilder {
	b.server.WriteTimeout = writeTimeout
	return b
}

func (b *ServerBuilder) Middleware(middleware ...string) *ServerBuilder {
	b.server.Middleware = append(b.server.Middleware, middleware...)
	return b
}

func (b *ServerBuilder) TLS(tls bool) *ServerBuilder {
	b.server.TLS = tls
	return b
}

func (b *ServerBuilder) Reset() *ServerBuilder {
	b.server = &Server{
		Port: 8080,
	}
	return b
}

func (b *ServerBuilder) Build() (*Server, error) {
	if b.server.Port <= 0 {
		return nil, errors.New("port must be positive")
	}
	if b.server.ReadTimeout < 0 {
		return nil, errors.New("read timeout cannot be negative")
	}
	if b.server.WriteTimeout < 0 {
		return nil, errors.New("write timeout cannot be negative")
	}

	result := &Server{
		Port:         b.server.Port,
		ReadTimeout:  b.server.ReadTimeout,
		WriteTimeout: b.server.WriteTimeout,
		Middleware:   append([]string{}, b.server.Middleware...),
		TLS:          b.server.TLS,
	}

	b.Reset()

	return result, nil
}

func (c *Server) String() string {
	var parts []string
	parts = append(parts, fmt.Sprintf("Port: %d", c.Port))
	if c.ReadTimeout > 0 {
		parts = append(parts, fmt.Sprintf("ReadTimeout: %s", c.ReadTimeout))
	}
	if c.WriteTimeout > 0 {
		parts = append(parts, fmt.Sprintf("WriteTimeout: %s", c.WriteTimeout))
	}
	if len(c.Middleware) > 0 {
		parts = append(parts, fmt.Sprintf("Middleware: [%s]", strings.Join(c.Middleware, ", ")))
	}
	parts = append(parts, fmt.Sprintf("TLS: %t", c.TLS))
	return "ServerConfig{" + strings.Join(parts, ", ") + "}"
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{
		server: &Server{},
	}
}
