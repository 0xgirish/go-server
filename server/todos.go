package server

import (
	"github.com/0xgirish/go-server/repository"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetTodos(c *fiber.Ctx) error {
	todos, err := s.todos.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "cannot get todos",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}

func (s *Server) CreateTodo(c *fiber.Ctx) error {
	// parse post body to Todo struct
	var todo repository.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "cannot parse body",
		})
	}

	id, err := s.todos.Create(todo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "cannot create todo",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "todo created successfully",
		"id":      id,
	})
}
