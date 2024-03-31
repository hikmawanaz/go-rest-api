package task

import "time"

type CompleteSingleTaskInput struct {
	TaskId      uint      `json:"taskId" binding:"required"`
	CompletedAt time.Time `json:"completedAt" binding:"required"`
}
