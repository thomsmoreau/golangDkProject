package server_test

import (
	"net/http"
	"testing"

	"github.com/matryer/is"

	"golangdk/integrationtest"
)

func TestServer_Start(t *testing.T) {
	integrationtest.SkipIfShort(t)

	t.Run("starts the server and listens for requests", func(t *testing.T) {
		is := is.New(t)

		cleanup := integrationtest.CreateServer()
		defer cleanup()

		resp, err := http.Get("http://localhost:8081/")
		is.NoErr(err)
		// It will return 404 since no routes are setted up in CreateServer func from integration test
		// Once done the cleanup func will stop server
		is.Equal(http.StatusNotFound, resp.StatusCode)
	})
}
