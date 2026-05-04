package service

import (
	"context"
	"errors"
	"sharing-app/internal/model"
	"sharing-app/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ListUsers(ctx context.Context) ([]model.User, error) {
	return s.repo.GetAll(ctx)
}

func (s *UserService) CreateUser(ctx context.Context, fullname string, phone string) error {
	if fullname == "" {
		return errors.New("fullname cannot be empty")
	}

	if phone == "" {
		return errors.New("phone cannot be empty")
	}

	user := model.User{
		Fullname: fullname,
		Phone:    phone,
	}

	return s.repo.Create(ctx, user)
}

func (s *UserService) UpdateUser(
	ctx context.Context,
	id int64,
	fullname string,
	phone string,
) error {

	if id <= 0 {
		return errors.New("invalid user id")
	}

	user := model.User{
		ID:       id,
		Fullname: fullname,
		Phone:    phone,
	}

	return s.repo.Update(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid user id")
	}

	return s.repo.Delete(ctx, id)
}

func (s *UserService) GetTripStats(ctx context.Context) ([]model.UserTripStats, error) {
	return s.repo.GetTripStats(ctx)
}
