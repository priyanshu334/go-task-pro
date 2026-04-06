package task

import (
	"time"

	"github.com/priyanshu334/taskmanage2/internal/pkg/utils"
)

type Status string
type Priority string

const (
	StatusTodo  Status = "todo"
	StatusDoing Status = "doing"
	StatusDone  Status = "done"

	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

type Task struct {
	utils.BaseModel

	Title       string `gorm:"not null"`
	Description string
	Status      Status   `gorm:"type:varchar(20);default:'todo'"`
	Priority    Priority `gorm:"type:varchar(20);default:'medium'"`
	DueDate     *time.Time
	UserID      string `gorm:"index"`
}
