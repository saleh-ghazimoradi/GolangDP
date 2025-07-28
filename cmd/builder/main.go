package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/builder"
	"time"
)

func main() {
	buick := builder.NewCarBuilder().
		Make("Buick").
		Model("Skylark").
		Color("white").
		Seats(5).
		Engine("V8").
		Transmission("Automatic").
		Build()
	fmt.Printf("%+v\n", buick)

	dbConfig := builder.NewDBConfigBuilder().
		Host("localhost").
		Port(5432).
		User("admin").
		Pass("secret").
		DbName("myapp").
		SslMode("require").
		MaxConn(20).
		DBBuild()
	fmt.Printf("%+v\n", dbConfig)

	query := builder.NewSQLQueryBuilder().
		Select("id", "name", "email").
		From("user").
		Where("age > 18").
		OrderBy("name ASC").
		Limit(10).
		Offset(20).
		Build()
	fmt.Printf("%+v\n", query)

	s := builder.NewServerBuilder()
	config, err := s.
		Port(9000).
		ReadTimeout(5*time.Second).
		WriteTimeout(10*time.Second).
		Middleware("auth", "logging").
		TLS(true).
		Build()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(config.String())
}
