package task

import "github.com/priyanshu334/taskmanage2/internal/database"

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(task *Task) error {
	return database.DB.Create(task).Error
}

func (r *Repository) FindAll(userID string, status string, search string, limit int, offset int) ([]Task, error) {
	var tasks []Task

	query := database.DB.Where("user_id = ?", userID)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if search != "" {
		query = query.Where("title ILIKE ?", "%"+search+"%")
	}

	err := query.Limit(limit).Offset(offset).Find(&tasks).Error
	return tasks, err
}

func (r *Repository) FindByID(id string) (*Task, error) {
	var task Task
	err := database.DB.First(&task, "id=?", id).Error
	return &task, err
}

func (r *Repository) Update(task *Task) error {
	return database.DB.Save(task).Error
}

func (r *Repository) Delete(task *Task) error {
	return database.DB.Delete(task).Error
}
