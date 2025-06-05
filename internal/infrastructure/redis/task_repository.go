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

func (r *TaskRepository) CreateTaskCallback(title string) (*task.Task, error) {

	id, err := r.client.Incr(r.ctx, "task:counter").Result()
	if err != nil {
		return nil, err
	}

	newTask := &task.Task{
		ID:    int(id),
		Title: title,
	}

	taskData, err := json.Marshal(newTask)
	if err != nil {
		return nil, err
	}

	err = r.client.Set(r.ctx, "task:"+string(rune(id)), taskData, 0).Err()
	if err != nil {
		return nil, err
	}

	err = r.client.LPush(r.ctx, "tasks:list", string(rune(id))).Err()
	if err != nil {
		return nil, err
	}

	return newTask, nil
}
