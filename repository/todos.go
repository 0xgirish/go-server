package repository

import "github.com/jmoiron/sqlx"

type Todo struct {
	ID        int64  `db:"id"        json:"id"`
	Task      string `db:"task"      json:"task"`
	Completed bool   `db:"completed" json:"completed"`
}

type Todos interface {
	GetAll() ([]Todo, error)
	GetByID(id int) (Todo, error)
	Create(todo Todo) (int64, error)
	MarkAsCompleted(id int64) error
}

type todos struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) Todos {
	return &todos{db: db}
}

func (t *todos) GetAll() ([]Todo, error) {
	query := `SELECT * FROM todos`

	var todos []Todo
	if err := t.db.Select(&todos, query); err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *todos) GetByID(id int) (Todo, error) {
	query := `SELECT * FROM todos WHERE id = ?`

	var todo Todo
	if err := t.db.Get(&todo, query, id); err != nil {
		return Todo{}, err
	}

	return todo, nil
}

func (t *todos) Create(todo Todo) (int64, error) {
	query := `INSERT INTO todos (task, completed) VALUES (:task, :completed)`

	result, err := t.db.NamedExec(query, todo)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (t *todos) MarkAsCompleted(id int64) error {
	query := `UPDATE todos SET completed = true WHERE id = ?`

	if _, err := t.db.Exec(query, id); err != nil {
		return err
	}

	return nil
}
