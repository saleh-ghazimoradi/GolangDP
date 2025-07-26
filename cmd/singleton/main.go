package main

import (
	"context"
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/singleton"
	"log/slog"
)

func main() {
	s := singleton.GetInstance()
	s.SetData("Hello World")
	fmt.Println(s.GetData())

	// Get the singleton Config instance
	config := singleton.GetConfig()
	fmt.Printf("Database URL: %s\n", config.DatabaseURL)
	fmt.Printf("API Key: %s\n", config.ApiKey)

	// Get the singleton DBPool instance
	dbPool := singleton.GetDBPool()
	err := dbPool.Db.PingContext(context.Background())
	if err != nil {
		fmt.Printf("Database connection failed: %v\n", err)
		return
	}
	fmt.Println("Database connection successful")
	
	// Get the singleton Logger instance
	logger := singleton.GetLogger()
	logger.Log.Info("Application started successfully", slog.String("app", "my-app"))

}
