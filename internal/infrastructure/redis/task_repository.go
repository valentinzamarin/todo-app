package redis

import (
	"context"
	"encoding/json"
	"todo-app/internal/domain/task"

	"github.com/redis/go-redis/v9"
)

type TaskRepository struct {
	client *redis.Client
	ctx    context.Context
}

func NewTaskRepository(client *redis.Client) *TaskRepository {
	return &TaskRepository{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *TaskRepository) GetAll() ([]task.Task, error) {
	taskIDs, err := r.client.LRange(r.ctx, "tasks:list", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var tasks []task.Task
	for _, idStr := range taskIDs {
		taskData, err := r.client.Get(r.ctx, "task:"+idStr).Result()
		if err != nil {
			continue
		}

		var t task.Task
		if err := json.Unmarshal([]byte(taskData), &t); err != nil {
			continue
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}
