package task

import (
	"time"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(req *CreateTaskRequest, userID string) error {
	var due *time.Time
	if req.DueDate != "" {
		t, _ := time.Parse(time.RFC3339, req.DueDate)
		due = &t
	}

	task := &Task{
		Title:       req.Title,
		Description: req.Description,
		Priority:    Priority(req.Priority),
		DueDate:     due,
		UserID:      userID,
	}
	return s.repo.Create(task)
}

func (s *Service) GetAll(userID, status, search string, limit, offset int) ([]Task, error) {
	return s.repo.FindAll(userID, status, search, limit, offset)
}
