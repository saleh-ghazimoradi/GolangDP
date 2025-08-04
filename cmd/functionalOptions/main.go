package main

import (
	"github.com/saleh-ghazimoradi/GolangDP/functionalOptions"
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
}
