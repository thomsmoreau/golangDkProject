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
		// Handler set for "/", should receive a 200 status code
		is.Equal(http.StatusOK, resp.StatusCode)
	})
}
