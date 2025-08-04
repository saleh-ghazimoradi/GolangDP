package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/functionalOptions"
	"os"
	"time"
)

func main() {
	server1 := functionalOptions.NewServer()
	server1.Start()
	server2 := functionalOptions.NewServer(
		functionalOptions.WithPort(9090),
		functionalOptions.WithTimeout(1*time.Second),
	)
	server2.Start()
	server3 := functionalOptions.NewServer(
		functionalOptions.WithPort(443),
		functionalOptions.WithTLSEnabled("server.crt", "server.key"),
	)
	server3.Start()

	fmt.Println()

	client1 := functionalOptions.NewDBClient()
	client1.Connect()

	client2 := functionalOptions.NewDBClient(
		functionalOptions.WithDBHost("db.example.com"),
		functionalOptions.WithDBPort(3306),
	)
	client2.Connect()

	client3 := functionalOptions.NewDBClient(
		functionalOptions.WithDBMaxPoolSize(50),
		functionalOptions.WithDBTimeout(10*time.Second),
		functionalOptions.WithDBRetryAttempts(5),
	)
	client3.Connect()
	fmt.Println()

	logger1 := functionalOptions.NewLogger()
	logger1.Log("This is a test message")
	logger2 := functionalOptions.NewLogger(
		functionalOptions.WithLevel("DEBUG"),
		functionalOptions.WithPrefix("APP1: "),
	)
	logger2.Log("Debugging the application")

	logger3 := functionalOptions.NewLogger(
		functionalOptions.WithLevel("ERROR"),
		functionalOptions.WithOutput(os.Stderr),
		functionalOptions.WithPrefix("ERROR: "),
	)

	logger3.Log("An error occurred")
}
