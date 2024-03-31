package main

import (
	"belyfe/auth"
	"belyfe/handler"
	"belyfe/helper"
	"belyfe/task"
	"belyfe/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=xbie password= dbname=belyfe port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	//Users module
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	taskRepository := task.NewRepository(db)
	taskService := task.NewService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService, authService)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	api := router.Group("api/v1")

	//Auth and users path
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", helper.AuthMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/users/fetch", helper.AuthMiddleware(authService, userService), userHandler.FetchUser)
	api.POST("/user/update/:id", helper.AuthMiddleware(authService, userService), userHandler.UpdateUser)

	api.POST("/tasks", helper.AuthMiddleware(authService, userService), taskHandler.CreateTask)
	api.GET("/tasks", helper.AuthMiddleware(authService, userService), taskHandler.FindAllTasks)
	api.POST("/tasks/complete-single-task", helper.AuthMiddleware(authService, userService), taskHandler.CompleteSingleTask)
	router.Run(":8081")
}
