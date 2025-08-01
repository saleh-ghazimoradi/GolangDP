package decorator

import (
	"fmt"
	"net/http"
	"time"
)

// This is the structural example of decorator pattern

// Handler interface defines the contract for handling HTTP requests
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type CoreHandler struct{}

func (h *CoreHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

type LoggerDecorator struct {
	Handler Handler
}

func (d *LoggerDecorator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request received: %s %s\n", r.Method, r.URL.Path)
	d.Handler.ServeHTTP(w, r)
	fmt.Println("Request processed")
}

type AuthDecorator struct {
	Handler Handler
}

func (d *AuthDecorator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") == "valid-token" {
		d.Handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

// This is the example of the functional decorator pattern

type QueryFunc func() (string, error)

func CoreQuery() (string, error) {
	time.Sleep(100 * time.Millisecond)
	return "Query result", nil
}

func TimingDecorator(query QueryFunc) QueryFunc {
	return func() (string, error) {
		start := time.Now()
		res, err := query()
		elapsed := time.Since(start)
		fmt.Printf("query took %v\n", elapsed)
		return res, err
	}
}

func RetryDecorator(query QueryFunc, maxAttempts int) QueryFunc {
	return func() (string, error) {
		for attempt := 1; attempt <= maxAttempts; attempt++ {
			res, err := query()
			if err == nil {
				return res, nil
			}
			fmt.Printf("Attempt %d failed: %v\n", attempt, err)
			if attempt < maxAttempts {
				time.Sleep(time.Duration(attempt) * 100 * time.Millisecond)
			}
		}
		return "", fmt.Errorf("failed after %d attempts", maxAttempts)
	}
}
