package integrationtest

import (
	"net/http"
	"testing"
	"time"

	"golangdk/server"
)

// CreateServer for testing on port 8081, returning a cleanup function that stops the server.
// Usage:
//
//	cleanup := CreateServer()
//	defer cleanup()
func CreateServer() func() {
	// Create new server struct
	s := server.New(server.Options{
		Host: "localhost",
		Port: 8081,
	})

	// Goroutine to panic if an error occurs with server
	go func() {
		if err := s.Start(); err != nil {
			panic(err)
		}
	}()

	// Wait for server to start
	for {
		_, err := http.Get("http://localhost:8081/")
		if err == nil {
			break
		}
		time.Sleep(5 * time.Millisecond)
	}

	// Return func to stop server
	return func() {
		if err := s.Stop(); err != nil {
			panic(err)
		}
	}
}

// SkipIfShort skips t if the "-short" flag is passed to "go test".
func SkipIfShort(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
}
