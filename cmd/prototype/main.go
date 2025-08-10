package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/prototype"
)

func main() {
	original := &prototype.Prototype{Name: "John", Age: 30}
	clone := original.Clone().(*prototype.Prototype)
	fmt.Println(original)
	fmt.Println(clone)
	fmt.Println()

	webConfigTemplate := &prototype.WebServerConfig{
		Port:        8080,
		MaxClients:  100,
		Timeout:     30,
		StaticFiles: []string{"index.html", "styles.css"},
	}

	dbConfigTemplate := &prototype.DBServerConfig{
		Port:       5432,
		MaxConn:    50,
		DBType:     "PostgreSQL",
		BackupFreq: 24,
	}

	webConfig1 := webConfigTemplate.Clone().(*prototype.WebServerConfig)
	webConfig1.Port = 80
	webConfig1.StaticFiles = append(webConfig1.StaticFiles, "app.js")

	dbConfig1 := dbConfigTemplate.Clone().(*prototype.DBServerConfig)
	dbConfig1.MaxConn = 100
	dbConfig1.BackupFreq = 12

	webConfig1.Display()
	dbConfig1.Display()

	webConfigTemplate.Display()
	dbConfigTemplate.Display()

}
