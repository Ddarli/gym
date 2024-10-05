package services

import (
	"context"
	"errors"
	"fmt"
	mylogger "github.com/Ddarli/gym/common/logger"
	"github.com/Ddarli/gym/userservice/models"
	"github.com/Ddarli/gym/userservice/repository"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"time"
)

type Service struct {
	models.UnimplementedUserServiceServer
	repo   repository.UserRepository
	logger *zap.SugaredLogger
}

func NewUserService(repo repository.UserRepository) models.UserServiceServer {
	return &Service{
		repo:   repo,
		logger: mylogger.GetLogger(),
	}
}

func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      expirationTime,
	})
	tokenString, err := token.SignedString([]byte("9b2f1c387a2f4ec9b6fa1f0481dfd3e5c56debc7e57c82b48cdbe8438f9e4a2d"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
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
func (s *Service) UpdateUser(ctx context.Context, in *models.UpdateUserRequest) (*models.User, error) {
	user := in.GetUser()
	err := s.repo.Update(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) Authenticate(ctx context.Context, in *models.AuthenticateRequest) (*models.AuthenticateResponse, error) {
	user, err := s.repo.GetByUsername(in.Username)
	response := models.AuthenticateResponse{}
	if err != nil {
		response.Error = "No user with such username"
		s.logger.Errorf("No user with such username")
		return &response, err
	}
	if in.Username == user.Username && in.Password == user.Password {
		token, err := generateToken(in.Username)
		if err != nil {
			s.logger.Errorf("Failed to generate token")
			return &response, err
		}
		response.Token = token
		return &response, nil
	}
	response.Error = "Invalid username or password"
	s.logger.Errorf("Invalid username or password")
	return &response, nil
}

func (s *Service) VerifyToken(ctx context.Context, in *models.VerifyTokenRequest) (*models.VerifyTokenResponse, error) {
	requestToken := in.GetToken()
	if requestToken == "" {
		s.logger.Errorf("Invalid token")
		return &models.VerifyTokenResponse{}, errors.New("empty token")
	}
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("9b2f1c387a2f4ec9b6fa1f0481dfd3e5c56debc7e57c82b48cdbe8438f9e4a2d"), nil
	})
	if err != nil {
		s.logger.Errorf("Invalid token: %v", err)
		return &models.VerifyTokenResponse{Result: false}, errors.New("invalid token")
	}
	response := &models.VerifyTokenResponse{Result: true}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				response.Result = false
				s.logger.Errorf("Token has expired")
				return response, nil
			}
		}
		return response, nil
	}
	return nil, nil
}
