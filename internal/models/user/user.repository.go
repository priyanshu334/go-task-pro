package user

import "github.com/priyanshu334/taskmanage2/internal/database"

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(user *User) error {
	return database.DB.Create(user).Error

}

func (r *Repository) FindByEmail(email string) (*User, error) {
	var user User
	err := database.DB.Where("email=?", email).First(&user).Error
	return &user, err
}
