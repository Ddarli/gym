package repository

import (
	"github.com/Ddarli/gym/classservice/models"
	"github.com/Ddarli/gym/common/logger"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type postgresClassRepository struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewPostgresClassRepository(db *sqlx.DB) ClassRepository {
	return &postgresClassRepository{
		db:     db,
		logger: logger.GetLogger(),
	}
}

func (r *postgresClassRepository) Create(class *models.ClassModel) (*models.ClassModel, error) {
	tx := r.db.MustBegin()
	tx.MustExec("INSERT INTO classes(name, description, capacity) VALUES($1, $2, $3)",
		class.Name, class.Description, class.Capacity)
	err := tx.Commit()
	if err != nil {
		r.logger.Errorf("create new class error: %s", err.Error())
		return nil, err
	}
	r.logger.Infof("create new class success")
	return class, nil
}
func (r *postgresClassRepository) Get(id int) (*models.ClassModel, error) {
	var class models.ClassModel
	err := r.db.Get(&class, "SELECT * FROM classes WHERE id = $1", id)
	if err != nil {
		r.logger.Errorf("get class error: %v", err.Error())
		return nil, err
	}
	r.logger.Infof("get class success")
	return &class, nil
}

func (r *postgresClassRepository) GetAll() ([]*models.ClassModel, error) {
	var classes []*models.ClassModel
	err := r.db.Select(&classes, "SELECT * FROM classes")
	if err != nil {
		r.logger.Errorf("get all classes error: %v", err.Error())
	}
	r.logger.Infof("get all classes success")
	return classes, nil
}
