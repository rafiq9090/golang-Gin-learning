package cache

import (
	"go_project_Gin/internal/model"
	"sync"
	"time"
)

type TaskCache struct {
	tasks  map[uint][]model.Task
	mu     sync.RWMutex
	expiry time.Time
}

var taskCache = &TaskCache{
	tasks: make(map[uint][]model.Task),
}

func GetTasks(userID uint) ([]model.Task, bool) {
	taskCache.mu.RLock()
	tasks, found := taskCache.tasks[userID]
	expired := time.Now().After(taskCache.expiry)
	taskCache.mu.RUnlock()

	if found && !expired {
		return tasks, true
	}
	return nil, false
}

func SetTasks(userID uint, tasks []model.Task) {
	taskCache.mu.Lock()
	taskCache.tasks[userID] = tasks
	taskCache.expiry = time.Now().Add(5 * time.Minute)
	taskCache.mu.Unlock()

	time.AfterFunc(time.Until(taskCache.expiry), func() {
		taskCache.mu.Lock()
		delete(taskCache.tasks, userID)
		taskCache.mu.Unlock()
	})
}
