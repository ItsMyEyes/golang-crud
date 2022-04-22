package routes

import (
	"crud_v2/handler"
	"crud_v2/middleware"
	. "crud_v2/utils"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	if data := os.Getenv("ENV"); data == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	gin.DisableConsoleColor()

	v1 := r.Group("/api/v1")

	anon := v1.Group("/").Use(middleware.AuthMiddleware)
	home := v1.Group("/home")
	todo := v1.Group("/todo").Use(middleware.AuthMiddleware)
	user := v1.Group("user")

	r.NoRoute(func(c *gin.Context) {
		message := map[string]string{"message": "Page is not found / not found method"}
		ResponseJson(c, http.StatusNotFound, message)
	})

	home.GET("/", handler.HomeHandler)

	todo.GET("/", handler.GetTodo)
	todo.POST("/", handler.PostTodo)
	todo.GET("/:id", handler.GetTodoById)
	todo.DELETE("/:id", handler.DeleteTodoById)

	user.POST("/register", handler.CreateUser)
	user.POST("/login", handler.LoginHandler)
	anon.GET("user/logout", handler.LogoutHandler)

	anon.GET("/", handler.HomePage)

	return r
}

func RunApplication() {
	var PORT string = "8080"
	if data := os.Getenv("PORT"); data != "" {
		PORT = data
	}
	router := SetupRoutes()
	router.SetTrustedProxies([]string{"127.0.0.1", "::1"})
	router.TrustedPlatform = "X-CDN-IP"

	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), router))
	router.Run(fmt.Sprintf(":%s", PORT))
}
