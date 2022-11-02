package server

import (
	"golangdk/handlers"
)

func (s *Server) setupRoutes() {
	handlers.Health(s.mux)
	handlers.FrontPage(s.mux)
}
