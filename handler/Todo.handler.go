package handler

import (
	"crud_v2/app/database"
	"crud_v2/entity"
	"crud_v2/model"
	"crud_v2/repository"
	. "crud_v2/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func GetTodo(c *gin.Context) {
	pagination := repository.PaginateTodo(c, &model.Pagination{})
	ResponseJson(c, http.StatusOK, pagination)
}

func PostTodo(c *gin.Context) {
	var todo entity.Todo
	if err := c.ShouldBind(&todo); err != nil {
		ResponseJson(c, http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if todo.Title == "" {
		ResponseJson(c, http.StatusBadRequest, gin.H{
			"error": "Title is required",
		})
		return
	}
	todo.Completed = false
	todo.OwnerId = uint(c.MustGet("user").(entity.User).Id)
	todo.User = c.MustGet("user").(entity.User)
	database.Connector.Create(&todo)
	ResponseJson(c, http.StatusOK, todo)
}

func GetTodoById(c *gin.Context) {
	id := c.Param("id")
	todo := repository.FindTodoById(id)
	log.Println("Iam")
	if todo.ID == 0 {
		ResponseJson(c, http.StatusNotFound, "Todo not found")
		return
	}
	ResponseJson(c, http.StatusOK, todo)
}

func DeleteTodoById(c *gin.Context) {
	id := c.Param("id")
	todo := repository.FindTodoById(id)
	if todo.ID == 0 {
		ResponseJson(c, http.StatusNotFound, "Todo not found")
		return
	}
	repository.DeleteTodoById(id)
	ResponseJson(c, http.StatusOK, "Todo deleted")
}
