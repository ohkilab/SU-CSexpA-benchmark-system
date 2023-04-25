package task

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client interface {
	SetTask(ctx context.Context, key string, task *Task) error
}

type client struct {
	rds *redis.Client
}

func NewClient(rds *redis.Client) Client {
	return &client{rds}
}

func (c *client) SetTask(ctx context.Context, key string, task *Task) error {
	return c.rds.Set(ctx, key, task, time.Hour).Err()
}

type TaskStatus string

const (
	TaskStatusWait = "wait"
	TaskStatusDone = "done"
)

type Task struct {
	IPAddr    string
	GroupID   string
	CreatedAt time.Time
}

func NewTask(ipAddr, groupID string) (*Task, string) {
	task := &Task{
		IPAddr:    ipAddr,
		GroupID:   groupID,
		CreatedAt: time.Now(),
	}
	return task, NewTaskKey(task, TaskStatusWait)
}

func NewTaskKey(task *Task, status TaskStatus) string {
	return fmt.Sprintf("%v_%v_%v_%v", task.GroupID, task.IPAddr, task.CreatedAt.UnixMilli(), status)
}
