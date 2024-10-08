package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Ddarli/gym/common/logger"
	"github.com/Ddarli/gym/shceduleservice/models"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type postgresScheduleRepository struct {
	logger *zap.SugaredLogger
	db     *sqlx.DB
}

func NewPostgresScheduleRepository(db *sqlx.DB) ScheduleRepository {
	return &postgresScheduleRepository{
		logger: logger.GetLogger(),
		db:     db,
	}
}

func (r *postgresScheduleRepository) GetById(id int) (models.ScheduleModel, error) {
	model := models.ScheduleModel{}
	err := r.db.Get(&model, "SELECT * FROM schedules WHERE id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.logger.Errorf("No model found for id: %d", id)
			return model, fmt.Errorf("schedule with id %d not found", id)
		}
		r.logger.Errorf("Error fetching model by id: %s", err.Error())
		return model, err
	}
	r.logger.Infof("Successfully fetched model by id: %d", id)
	return model, nil

}
func (r *postgresScheduleRepository) Create(model models.ScheduleModel) error {
	tx := r.db.MustBegin()
	tx.MustExec("INSERT INTO schedules (class_id, trainer_id, start_time, end_time) VALUES ($1, $2, $3, $4)",
		model.ClassId, model.TrainerId, model.StartTime, model.EndTime)
	err := tx.Commit()
	if err != nil {
		r.logger.Errorf("Error committing transaction: %s", err.Error())
		return err
	}
	return nil
}
