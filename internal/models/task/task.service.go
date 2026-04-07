package task

import (
	"errors"
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

func (s *Service) Update(taskID string, userID string, req *UpdateTaskRequest) error {
	task, err := s.repo.FindByID(taskID)
	if err != nil {
		return err
	}
	if task.UserID != userID {
		return errors.New("unautorized")
	}
	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Status != nil {
		task.Status = Status(*req.Status)
	}
	if req.Priority != nil {
		task.Priority = Priority(*req.Priority)
	}
	return s.repo.Update(task)

}

func (s *Service) Delete(taskID string, userID string) error {
	task, err := s.repo.FindByID(taskID)
	if err != nil {
		return err
	}
	if task.UserID != userID {
		return errors.New("unautorized")
	}
	return s.repo.Delete(task)
}
