package server

import (
	"context"
	"net/http"
	"time"
)

// Server - type of http server
type Server struct {
	httpServer *http.Server
}

// Run - run server
func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

// Shutdown - stop server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
