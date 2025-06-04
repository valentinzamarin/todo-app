package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-app/internal/app/task"
	"todo-app/internal/infrastructure/config"
	"todo-app/internal/infrastructure/redis"
	"todo-app/internal/interfaces/http/handlers"

	server "todo-app/internal/interfaces/http"

	goredis "github.com/redis/go-redis/v9"
)

func main() {

	cfg := config.Load()

	redisClient := goredis.NewClient(&goredis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
	})

	taskRepo := redis.NewTaskRepository(redisClient)

	getTasksUseCase := task.NewGetTasksUseCase(taskRepo)

	taskHandlers := handlers.NewTaskHandler(getTasksUseCase)

	r := server.NewRouter(taskHandlers)

	serverAddr := ":" + cfg.ServerPort
	log.Fatal(http.ListenAndServe(serverAddr, r))
}
