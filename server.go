package booksAndAuthors

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) RunServer(port string, router *http.ServeMux) error {
	s.httpServer = &http.Server{Addr: port,
		Handler:      router,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutdownServer(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
