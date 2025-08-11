package abstractFactory

import "fmt"

type DBConnection interface {
	Connect() string
}

type DBQuery interface {
	Execute() string
}

type MySQLConnection struct{}

func (m *MySQLConnection) Connect() string {
	return "Connected to MySQL"
}

type MySQLQuery struct{}

func (m *MySQLQuery) Execute() string {
	return "Executed MySQL Query"
}

type PostgreSQLConnection struct{}

func (p *PostgreSQLConnection) Connect() string {
	return "Connected to PostgreSQL"
}

type PostgreSQLQuery struct{}

func (p *PostgreSQLQuery) Execute() string {
	return "Executed PostgreSQL Query"
}

type DBFactory interface {
	CreateConnection() DBConnection
	CreateQuery() DBQuery
}

type MySQLFactory struct{}

func (m *MySQLFactory) CreateConnection() DBConnection {
	return &MySQLConnection{}
}

func (m *MySQLFactory) CreateQuery() DBQuery {
	return &MySQLQuery{}
}

type PostgreSQLFactory struct{}

func (p *PostgreSQLFactory) CreateConnection() DBConnection {
	return &PostgreSQLConnection{}
}

func (p *PostgreSQLFactory) CreateQuery() DBQuery {
	return &PostgreSQLQuery{}
}

func UseDatabase(factory DBFactory) {
	connection := factory.CreateConnection()
	query := factory.CreateQuery()
	fmt.Println(connection.Connect())
	fmt.Println(query.Execute())
}
