package main

import (
	"github.com/Ddarli/gym/common"
	"github.com/Ddarli/gym/trainerservice/db"
	"github.com/Ddarli/gym/trainerservice/models"
	"github.com/Ddarli/gym/trainerservice/repository"
	"github.com/Ddarli/gym/trainerservice/services"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	service := NewService()
	address := os.Getenv("ADDRESS")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	models.RegisterTrainerServiceServer(grpcServer, service)
	log.Printf("Serving gRPC on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}

func NewService() models.TrainerServiceServer {
	conf := common.LoadConfig()
	db, err := db.NewPostgresConnection(conf)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewPostgresTrainerRepository(db)
	trainerService := services.NewTrainerService(repo)
	return trainerService
}
