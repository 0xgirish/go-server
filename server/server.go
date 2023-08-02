package server

import (
	"fmt"

	"github.com/0xgirish/go-server/repository"
	"github.com/gofiber/fiber/v2"
)

// Server is a server, this server can register apis
// apis can decide how to register themselves
type Server struct {
	todos repository.Todos
	app   *fiber.App
}

// New returns a new Server instance
func New(todos repository.Todos) *Server {
	s := &Server{app: fiber.New(), todos: todos}
	s.register()

	return s
}

func (s *Server) register() {
	app := s.app

	// add your apis here
	app.Get("/hello", s.SayHello)
	app.Post("/todo/create", s.CreateTodo)
	app.Get("/todo/all", s.GetTodos)
}

// Start starts the server
func (s *Server) Start(host string, port int) error {
	return s.app.Listen(fmt.Sprintf("%s:%d", host, port))
}
