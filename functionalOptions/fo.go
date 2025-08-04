package functionalOptions

import (
	"fmt"
	"time"
)

/*
	The functional options pattern is a design pattern in Golang that provides a flexible and idiomatic way to configure objects with optional settings. Avoiding a large number of parameters or configuration struct.
*/

type Server struct {
	Port       int
	Timeout    time.Duration
	TLSEnabled bool
	CertFile   string
	KeyFile    string
}

type Option func(*Server)

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func WithTLSEnabled(certFile, keyFile string) Option {
	return func(s *Server) {
		s.TLSEnabled = true
		s.CertFile = certFile
		s.KeyFile = keyFile
	}
}

func (s *Server) Start() {
	fmt.Printf("Starting server on port %d with timeout %v\n", s.Port, s.Timeout)
	if s.TLSEnabled {
		fmt.Printf("TLS enabled with cert: %s, key: %s\n", s.CertFile, s.KeyFile)
	} else {
		fmt.Println("TLS disabled")
	}
}

func NewServer(opts ...Option) *Server {
	s := &Server{
		Port:    8080,
		Timeout: 10 * time.Second,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
