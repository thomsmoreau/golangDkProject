/*
Package server contains everything for setting up and running the HTTP server.
*/
package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Server struct {
	address string
	mux     chi.Router
	server  *http.Server
	log     *zap.Logger
}

type Options struct {
	Host string
	Log  *zap.Logger
	Port int
}

// New turns the host and port pair into an address the server understands (in the form host:port)
//
//	and creates a http.Server using that address and the mux.
func New(opts Options) *Server {
	// nil check for logger
	if opts.Log == nil {
		opts.Log = zap.NewNop()
	}

	// Join the host and port with a colon.
	address := net.JoinHostPort(opts.Host, strconv.Itoa(opts.Port))

	// Create a new router.
	mux := chi.NewMux()

	// Create a new server.
	return &Server{
		address: address,
		mux:     mux,
		log:     opts.Log,
		server: &http.Server{
			Addr:    address,
			Handler: mux,
			// Set up some timeouts
			ReadTimeout:       5 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      5 * time.Second,
			IdleTimeout:       5 * time.Second,
		},
	}
}

// Start the Server by setting up routes and listening for HTTP requests on the given address.
func (s *Server) Start() error {
	// Set up handlers mux
	s.setupRoutes()

	s.log.Info("Starting server ", zap.String("address", s.address))

	// ListenAndServe always return http.ErrServerClosed, check if different
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("error starting server: %w", err)
	}
	return nil
}

// Stop the Server gracefully within the timeout.
func (s *Server) Stop() error {
	s.log.Info("Stopping")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("error stopping server: %w", err)
	}

	return nil
}
