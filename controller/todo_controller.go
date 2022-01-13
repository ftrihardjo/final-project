package controller

import (
	"final-project/db"
	"final-project/model"
	"net/http"
	"strconv"

	"github.com/swaggo/swag/example/celler/httputil"

	"github.com/gin-gonic/gin"
)

type TodoAttribute struct {
	Title       string
	Description string
	Completed   bool
}

// CreateTodo godoc
// @Tags todos
// @Description creating a new todo
// @ID create-todo
// @Accept json
// @Param RequestBody body TodoAttribute true "request body json"
// @Router /todos [post]
func CreateTodo(c *gin.Context) {
	db := db.GetDB()
	todo := model.Todo{}
	if err := c.ShouldBindJSON(&todo); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	statement := "INSERT INTO todos (title,description,completed) VALUES ($1, $2,$3)"
	db.QueryRow(statement, todo.Title, todo.Description, todo.Completed)
}

// GetTodos godoc
// @Tags todos
// @Description get all todos
// @ID get-todos
// @Sucess (202) {array} model.Todo
// @Produce      json
// @Router /todos [get]
func GetTodos(c *gin.Context) {
	db := db.GetDB()
	statement := "SELECT * FROM todos"
	rows, err := db.Query(statement)
	defer rows.Close()
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	var todos []model.Todo
	for rows.Next() {
		todo := model.Todo{}
		rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed)
		todos = append(todos, todo)
	}
	c.JSON(http.StatusOK, todos)
}

// UpdateTodo godoc
// @Tags todos
// @Description update a todo
// @ID update-todo
// @Accept json
// @Param        id   path      int  true  "Account ID"
// @Param RequestBody body TodoAttribute true "request body json"
// @Router /todos/{id} [put]
func UpdateTodo(c *gin.Context) {
	db := db.GetDB()
	todo := model.Todo{}
	if err := c.ShouldBindJSON(&todo); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	todo.Id = id
	statement := "UPDATE todos SET title=$1, description=$2, completed=$3 WHERE id=$4"
	db.QueryRow(statement, todo.Title, todo.Description, todo.Completed, todo.Id)
}

// GetTodo godoc
// @Tags todos
// @Description get a todo
// @ID get-todo
// @Produce json
// @Param        id   path      int  true  "Account ID"
// @Router /todos/{id} [get]
func GetTodo(c *gin.Context) {
	db := db.GetDB()
	todo := model.Todo{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	statement := "SELECT * FROM todos WHERE id=$1"
	row := db.QueryRow(statement, id)
	row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed)
	c.JSON(http.StatusOK, todo)
}

// DeleteTodo godoc
// @Tags todos
// @Description delete a todo
// @ID get-todo
// @Param        id   path      int  true  "Account ID"
// @Router /todos/{id} [delete]
func DeleteTodo(c *gin.Context) {
	db := db.GetDB()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	statement := "DELETE FROM todos WHERE id=$1"
	db.QueryRow(statement, id)
}
