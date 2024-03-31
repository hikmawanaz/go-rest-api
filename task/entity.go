package task

import "time"

type Task struct {
	ID                uint   `gorm:"primaryKey"`
	Name              string `gorm:"not null"`
	Description       string
	DueDate           time.Time
	RecurringType     string          `gorm:"default:null"` // "daily", "weekly", "monthly"
	RecurringInterval uint            `gorm:"default:null"`
	StartDate         time.Time       `gorm:"default:null"`
	EndDate           time.Time       `gorm:"default:null"`
	CompletedTasks    []CompletedTask `gorm:"foreignKey:TaskID"`
	UserID            uint            `gorm:"not null"`
}

type CompletedTask struct {
	ID          uint      `gorm:"primaryKey"`
	TaskID      uint      `gorm:"not null"`
	CompletedAt time.Time `gorm:"not null"`
}
