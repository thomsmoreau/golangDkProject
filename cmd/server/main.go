package main

import (
	"context"
	"fmt"
	"golangdk/server"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

// release is set through the linker at build time, generally from a git sha.
// Used for logging and error reporting.
var release string

func main() {
	fmt.Println("Message main.go")
	os.Exit(start())
}

func start() int {
	// Get the value of LOG_ENV env var
	logEnv := getStringOrDefault("LOG_ENV", "development")

	// Create logger depending of env passed
	log, err := createLogger(logEnv)
	if err != nil {
		fmt.Println("Error setting up the logger:", err)
		return 1
	}

	// Add fields to logger
	log = log.With(zap.String("release", release), zap.String("env", logEnv))

	defer func() {
		// If we cannot sync, there's probably something wrong with outputting logs,
		// so we probably cannot write using fmt.Println either. So just ignore the error.
		_ = log.Sync()
	}()

	// Get host and port from env
	host := getStringOrDefault("HOST", "localhost")
	port := getIntOrDefault("PORT", 8080)

	s := server.New(server.Options{
		Host: host,
		Port: port,
		Log:  log,
	})

	// An error group is something that can run functions in goroutines, wait for the functions to finish, and return any errors.
	var eg errgroup.Group
	// NotifyContext function makes sure to cancel the returned context if it receives one of the signals we have asked for.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	eg.Go(func() error {
		// Blocks until the context before created is  done (receive a SIGTERM or SIGINT)
		<-ctx.Done()
		if err := s.Stop(); err != nil {
			log.Info("Error stopping server", zap.Error(err))
			return err
		}
		return nil
	})

	if err := s.Start(); err != nil {
		log.Info("Error starting server", zap.Error(err))
		return 1
	}
	// By calling eg.Wait on the error group from before, we wait for all functions passed to calls to eg.Go to finish
	if err := eg.Wait(); err != nil {
		return 1
	}
	return 0
}

func createLogger(env string) (*zap.Logger, error) {
	switch env {
	case "production":
		return zap.NewProduction()
	case "development":
		return zap.NewDevelopment()
	default:
		return zap.NewNop(), nil
	}
}

// getStringOrDefault returns the value of the env name passed or defaultV if not found
func getStringOrDefault(name, defaultV string) string {
	v, ok := os.LookupEnv(name)
	if !ok {
		return defaultV
	}
	return v
}

func getIntOrDefault(name string, defaultV int) int {
	v, ok := os.LookupEnv(name)
	if !ok {
		return defaultV
	}
	vAsInt, err := strconv.Atoi(v)
	if err != nil {
		return defaultV
	}
	return vAsInt
}
