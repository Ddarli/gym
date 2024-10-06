package services

import (
	"context"
	"github.com/Ddarli/gym/classservice/models"
	"github.com/Ddarli/gym/classservice/repository"
	"github.com/Ddarli/gym/common/logger"
	"go.uber.org/zap"
	"strconv"
)

type service struct {
	models.UnimplementedClassServiceServer
	db     repository.ClassRepository
	logger *zap.SugaredLogger
}

func ToProto(class *models.ClassModel) *models.Class {
	return &models.Class{
		Id:          strconv.Itoa(class.Id),
		Name:        class.Name,
		Description: class.Description,
		Capacity:    int32(class.Capacity),
	}
}

func NewClassService(db repository.ClassRepository) models.ClassServiceServer {
	return &service{db: db, logger: logger.GetLogger()}
}

func (s *service) GetClass(ctx context.Context, request *models.GetClassRequest) (*models.GetClassResponse, error) {
	id, err := strconv.Atoi(request.GetId())
	if err != nil {
		s.logger.Errorf("Error converting string to int: %v", err)
		return nil, err
	}
	class, err := s.db.Get(id)
	return &models.GetClassResponse{Class: ToProto(class)}, err
}
func (s *service) CreateClass(ctx context.Context, request *models.CreateClassRequest) (*models.CreateClassResponse, error) {
	class := &models.ClassModel{
		Name:        request.Name,
		Description: request.Description,
		Capacity:    int(request.Capacity),
	}
	savedClass, err := s.db.Create(class)
	if err != nil {
		s.logger.Errorf("Error creating class: %v", err)
		return nil, err
	}
	s.logger.Infof("Created class successfully")
	return &models.CreateClassResponse{Class: ToProto(savedClass)}, nil
}

func (s *service) GetClasses(ctx context.Context, request *models.GetClassesRequest) (*models.GetClassesResponse, error) {
	classes, err := s.db.GetAll()
	if err != nil {
		s.logger.Errorf("Error getting classes: %v", err)
		return nil, err
	}
	responseClasses := make([]*models.Class, 0, len(classes))
	for _, class := range classes {
		responseClasses = append(responseClasses, ToProto(class))
	}
	return &models.GetClassesResponse{Classes: responseClasses}, nil
}
