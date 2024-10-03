package services

import (
	"context"
	"github.com/Ddarli/gym/userservice/models"
	"github.com/Ddarli/gym/userservice/repository"
)

type Service struct {
	models.UnimplementedUserServiceServer
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) models.UserServiceServer {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(ctx context.Context, in *models.CreateUserRequest) (*models.User, error) {
	user := models.User{
		Username:    in.Username,
		Password:    in.Password,
		Email:       in.Email,
		PhoneNumber: in.PhoneNumber,
	}
	err := s.repo.Create(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (s *Service) GetUser(ctx context.Context, in *models.GetUserRequest) (*models.User, error) {
	user, err := s.repo.Get(in.GetId())
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *Service) UpdateUser(ctx context.Context, in *models.UpdateUserRequest) (*models.User, error) {
	user := in.GetUser()
	err := s.repo.Update(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *Service) DeleteUser(ctx context.Context, in *models.DeleteUserRequest) (*models.DeleteUserResponse, error) {
	err := s.repo.Delete(in.Id)
	response := &models.DeleteUserResponse{}
	if err != nil {
		response.Success = false
		return response, err
	}
	response.Success = true
	return response, nil
}
