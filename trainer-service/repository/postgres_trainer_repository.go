package repository

import (
	"github.com/Ddarli/gym/common/logger"
	"github.com/Ddarli/gym/trainerservice/models"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type postgresTrainerRepository struct {
	logger *zap.SugaredLogger
	db     *sqlx.DB
}

func NewPostgresTrainerRepository(db *sqlx.DB) TrainerRepository {
	return &postgresTrainerRepository{
		logger: logger.GetLogger(),
		db:     db,
	}
}

func (r *postgresTrainerRepository) Get(id int) (*models.TrainerModel, error) {
	var trainer models.TrainerModel
	err := r.db.Get(&trainer, "SELECT * FROM trainers WHERE id = $1", id)
	if err != nil {
		r.logger.Errorf("Get trainers id=%d, err=%s", id, err)
		return nil, err
	}
	r.logger.Infof("Get trainers id=%d, trainer=%+v", id, trainer)
	return &trainer, nil
}
func (r *postgresTrainerRepository) Create(model *models.TrainerModel) (*models.TrainerModel, error) {
	tx := r.db.MustBegin()
	tx.MustExec("INSERT INTO trainers (name, specialization) VALUES ($1, $2)",
		model.Name, model.Specialization)
	err := tx.Commit()
	if err != nil {
		r.logger.Errorf("Create trainers err=%s", err)
		return nil, err
	}
	r.logger.Infof("Create trainers trainer=%+v", model)
	return model, nil

}
