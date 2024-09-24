package server

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(handler http.Handler, port string) *Server {
	return &Server{
		server: &http.Server{
			Addr:    port,
			Handler: handler,
		},
	}
}

func (s *Server) Start() error {
	logrus.Printf("Server running on %v", s.server.Addr)
	return s.server.ListenAndServe()
}
