package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Ddarli/gym/userservice/models"
	"github.com/jmoiron/sqlx"
	"log"
	"strconv"
)

type postgresUserRepository struct {
	db *sqlx.DB
}

func NewPostgresUserRepository(db *sqlx.DB) UserRepository {
	return &postgresUserRepository{db: db}
}

func fromProto(u *models.User) *models.UserModel {
	id, _ := strconv.Atoi(u.Id)
	return &models.UserModel{
		Id:          id,
		Username:    u.Username,
		Password:    u.Password,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
	}
}

func toProto(u *models.UserModel) *models.User {
	return &models.User{
		Id:          strconv.Itoa(u.Id),
		Username:    u.Username,
		Password:    u.Password,
		Email:       u.Email,
		PhoneNumber: u.PhoneNumber,
	}
}

func (r *postgresUserRepository) Create(user *models.User) error {
	tx := r.db.MustBegin()
	tx.MustExec("INSERT INTO users (username, password, email, phone_number) VALUES ($1, $2, $3, $4)",
		user.Username, user.Password, user.Email, user.PhoneNumber)
	err := tx.Commit()

	if err != nil {
		return err
	}
	return nil
}
func (r *postgresUserRepository) Get(userId string) (*models.User, error) {
	var user models.UserModel
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}
	err = r.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user with id %d not found", id)
		}
		return nil, err
	}
	return toProto(&user), nil
}
func (r *postgresUserRepository) Update(user *models.User) error {
	tx := r.db.MustBegin()
	log.Println(user)
	userId, err := strconv.Atoi(user.Id)
	if err != nil {
		return err
	}
	tx.MustExec("UPDATE users SET username=$1, password=$2, email=$3, phone_number=$4 WHERE id = $5",
		user.Username, user.Password, user.Email, user.PhoneNumber, userId)
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
func (r *postgresUserRepository) Delete(id string) error {
	userId, err := strconv.Atoi(id)
	tx := r.db.MustBegin()
	tx.MustExec("DELETE FROM users WHERE id = $1", userId)
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
