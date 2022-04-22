package repository

import (
	"crud_v2/app/database"
	"crud_v2/entity"
	"crud_v2/model"
	. "crud_v2/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PaginateTodo(c *gin.Context, pagination *model.Pagination) *model.Pagination {
	var todos []entity.Todo

	limit, _ := strconv.Atoi(c.Query("limit"))
	pagination.Limit = limit

	page, _ := strconv.Atoi(c.Query("page"))
	pagination.Page = page

	sort := c.Query("sort")
	pagination.Sort = sort

	database.Connector.Scopes(Paginate(todos, pagination, database.Connector, "User", "users.id, users.username, users.email, users.phone, users.status")).Find(&todos)
	pagination.Data = todos
	return pagination
}

func FindTodoByTitle(title string) *entity.Todo {
	var todo entity.Todo
	database.Connector.Model(&todo).Select("id, title, completed").Where("title = ?", title).First(&todo)
	return &todo
}

func FindTodoById(id string) *entity.Todo {
	var todo entity.Todo
	database.Connector.Model(&todo).Where("id = ?", id).First(&todo)
	return &todo
}

func FindTodoByIdWithUser(id string) *entity.Todo {
	var todo entity.Todo
	database.Connector.Model(&todo).Where("id = ?", id).Preload("User").First(&todo)
	return &todo
}

func DeleteTodoById(id string) {
	var todo entity.Todo
	database.Connector.Model(&todo).Where("id = ?", id).Delete(&todo)
}
