package service

import (
	"context"
	"time"

	"github.com/Infamous003/ainyx/internal/models"
	"github.com/Infamous003/ainyx/internal/repository"
)

type User struct {
	repo *repository.User
}

func NewUser(repo *repository.User) *User {
	return &User{repo: repo}
}

func (s *User) GetUser(ctx context.Context, id int32) (*models.UserRead, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &models.UserRead{
		ID:   int(id),
		Name: user.Name,
		Dob:  models.Date(user.Dob),
		Age:  CalculateAge(user.Dob),
	}, nil
}

func (s *User) ListUsers(ctx context.Context) ([]*models.UserRead, error) {
	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	res := []*models.UserRead{}
	for _, user := range users {
		res = append(res, &models.UserRead{
			ID:   int(user.ID),
			Name: user.Name,
			Dob:  models.Date(user.Dob),
			Age:  CalculateAge(user.Dob),
		})
	}

	return res, nil
}

func (s *User) CreateUser(ctx context.Context, name string, dob time.Time) (*models.UserBasic, error) {
	user, err := s.repo.Create(ctx, name, dob)
	if err != nil {
		return nil, err
	}

	return &models.UserBasic{
		ID:   int(user.ID),
		Name: user.Name,
		Dob:  models.Date(user.Dob),
	}, nil
}

func (s *User) UpdateUser(ctx context.Context, id int32, name string, dob time.Time) (*models.UserBasic, error) {
	user, err := s.repo.Update(ctx, id, name, dob)
	if err != nil {
		return nil, err
	}

	return &models.UserBasic{
		ID:   int(user.ID),
		Name: user.Name,
		Dob:  models.Date(user.Dob),
	}, nil
}

func (s *User) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.Delete(ctx, id)
}
