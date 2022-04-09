package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	Addr   string
}

func (s *Server) Initialize(addr string, r *mux.Router) {
	s.Router = r
	s.Addr = addr
}

func (s *Server) Run() {
	http.Handle("/", Cors.Handler(s.Router))
	log.Println("Server running on", s.Addr)
	log.Fatal(http.ListenAndServe(s.Addr, nil))
}

func NewServer() Server {
	return Server{}
}
