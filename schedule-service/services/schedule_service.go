package services

import (
	"context"
	"github.com/Ddarli/gym/common/logger"
	"github.com/Ddarli/gym/shceduleservice/models"
	"github.com/Ddarli/gym/shceduleservice/repository"
	"go.uber.org/zap"
	"strconv"
)

type service struct {
	models.UnimplementedScheduleServiceServer
	logger *zap.SugaredLogger
	repo   repository.ScheduleRepository
}

func NewScheduleService(repo repository.ScheduleRepository) models.ScheduleServiceServer {
	return &service{
		repo:   repo,
		logger: logger.GetLogger(),
	}
}

func (s *service) GetSchedule(ctx context.Context, in *models.GetScheduleRequest) (*models.GetScheduleResponse, error) {
	id, _ := strconv.Atoi(in.GetId())
	model, err := s.repo.GetById(id)
	if err != nil {
		s.logger.Error("ScheduleService GetSchedule error", zap.Error(err))
		return &models.GetScheduleResponse{Schedule: nil}, err
	}
	return &models.GetScheduleResponse{Schedule: models.ToProto(&model)}, nil
}
func (s *service) CreateSchedule(ctx context.Context, in *models.CreateScheduleRequest) (*models.CreateScheduleResponse, error) {
	model := models.ToModel(in)
	err := s.repo.Create(*model)
	if err != nil {
		s.logger.Error("ScheduleService CreateSchedule error", zap.Error(err))
		return nil, err
	}
	s.logger.Infof("ScheduleService CreateSchedule success")
	return &models.CreateScheduleResponse{Schedule: models.ToProto(model)}, nil
}
