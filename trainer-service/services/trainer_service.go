package services

import (
	"context"
	"github.com/Ddarli/gym/common/logger"
	"github.com/Ddarli/gym/trainerservice/models"
	"github.com/Ddarli/gym/trainerservice/repository"
	"go.uber.org/zap"
	"strconv"
)

type service struct {
	models.UnimplementedTrainerServiceServer
	repo   repository.TrainerRepository
	logger *zap.SugaredLogger
}

func NewTrainerService(repo repository.TrainerRepository) models.TrainerServiceServer {
	return &service{repo: repo, logger: logger.GetLogger()}
}

func (s *service) GetTrainer(ctx context.Context, req *models.GetTrainerRequest) (*models.GetTrainerResponse, error) {
	id, err := strconv.Atoi(req.GetId())
	if err != nil {
		s.logger.Warnw("get trainer id error", "id", req.GetId())
		return nil, err
	}
	trainer, err := s.repo.Get(id)
	if err != nil {
		s.logger.Warnw("get trainer error", "id", req.GetId())
		return nil, err
	}
	return &models.GetTrainerResponse{Trainer: models.ToProto(trainer)}, nil
}
func (s *service) CreateTrainer(ctx context.Context, req *models.CreateTrainerRequest) (*models.CreateTrainerResponse, error) {
	trainer, err := s.repo.Create(&models.TrainerModel{
		Name:           req.Name,
		Specialization: req.Specialization,
	})
	if err != nil {
		return nil, err
	}
	response := &models.CreateTrainerResponse{
		Trainer: models.ToProto(trainer),
	}
	return response, err
}
