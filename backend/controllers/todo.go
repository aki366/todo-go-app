package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "todo-app/models"
)

// todoListにTodoのリストを格納
var todoList = []models.TodoItem{}

// GetTodosはすべてのTodoをリストアップするGETリクエストを処理
func GetTodos(c *gin.Context) {
    c.JSON(http.StatusOK, todoList)
}

// AddTodoは、新しいTodoを追加するPOSTリクエストを処理
func AddTodo(c *gin.Context) {
    var newTodo models.TodoItem
    if err := c.BindJSON(&newTodo); err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }
    todoList = append(todoList, newTodo)
    c.Status(http.StatusCreated)
}

// DeleteTodo は、Todo を削除する DELETE リクエストを処理
func DeleteTodo(c *gin.Context) {
    id := c.Param("id")
    for i, item := range todoList {
        if item.ID == id {
            todoList = append(todoList[:i], todoList[i+1:]...)
            c.Status(http.StatusOK)
            return
        }
    }
    c.Status(http.StatusNotFound)
}
