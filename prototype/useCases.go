package prototype

import (
	"fmt"
	"time"
)

// USE CASE NUM 1

type ServerConfig interface {
	Clone() ServerConfig
	Display()
}

type WebServerConfig struct {
	Port        int
	MaxClients  int
	Timeout     time.Duration
	StaticFiles []string
}

func (w *WebServerConfig) Display() {
	fmt.Printf("WebServerConfig: Port=%d, MaxClients=%d, Timeout=%d, StaticFiles=%v\n",
		w.Port, w.MaxClients, w.Timeout, w.StaticFiles)
}

func (w *WebServerConfig) Clone() ServerConfig {
	staticFiles := make([]string, len(w.StaticFiles))
	copy(staticFiles, w.StaticFiles)
	return &WebServerConfig{
		Port:        w.Port,
		MaxClients:  w.MaxClients,
		Timeout:     w.Timeout,
		StaticFiles: staticFiles,
	}
}

type DBServerConfig struct {
	Port       int
	MaxConn    int
	DBType     string
	BackupFreq int
}

func (d *DBServerConfig) Clone() ServerConfig {
	return &DBServerConfig{
		Port:       d.Port,
		MaxConn:    d.MaxConn,
		DBType:     d.DBType,
		BackupFreq: d.BackupFreq,
	}
}

func (d *DBServerConfig) Display() {
	fmt.Printf("DBServerConfig: Port=%d, MaxConn=%d, DBType=%s, BackupFreq=%d\n",
		d.Port, d.MaxConn, d.DBType, d.BackupFreq)
}
