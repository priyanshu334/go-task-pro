package user

import (
	"errors"

	"github.com/priyanshu334/taskmanage2/internal/pkg/utils"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Register(req *RegisterRequest) error {
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := &User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashed,
	}

	return s.repo.Create(user)
}

func (s *Service) Login(req *LoginRequest) (*User, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := utils.CheckPassword(req.Password, user.Password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
