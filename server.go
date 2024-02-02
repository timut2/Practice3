package server

import (
	"net/http"
	"time"
	"context"
)
type Server struct{
	httpServer *http.Server
}

func (s *Server) Run(post string ) error{
	s.httpServer = &http.Server{
		Addr: ":"+port,
		MaxHeaderBytes: 1<<20,
		ReadTimeout:10 *
		ReadTimeout: 10 * time.Second
		WriteTimeout: 10 * time.Second

	}
	return s.httpServer.ListenAndServer()
}

func (s *Server) Shutdown(ctx context.Context) error{
	return s.httpServer.Shutdown(ctx)
}