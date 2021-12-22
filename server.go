package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) //Sintaxis que indica que no sabemos cuantos middlewares estaran involucrados

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	} else {
		return nil
	}
}
