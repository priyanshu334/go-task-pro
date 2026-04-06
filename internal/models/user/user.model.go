package user

import "github.com/priyanshu334/taskmanage2/internal/pkg/utils"

type User struct {
	utils.BaseModel

	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
