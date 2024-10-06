package main

import (
	"github.com/Ddarli/gym/common"
	"github.com/Ddarli/gym/userservice/db"
	"github.com/Ddarli/gym/userservice/models"
	"github.com/Ddarli/gym/userservice/repository"
	"github.com/Ddarli/gym/userservice/services"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func newService() models.UserServiceServer {
	conf := common.LoadConfig()
	db, err := db.NewPostgresConnection(conf)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewPostgresUserRepository(db)
	userService := services.NewUserService(repo)
	return userService
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	service := newService()
	address := os.Getenv("ADDRESS")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	models.RegisterUserServiceServer(grpcServer, service)
	log.Printf("Serving gRPC on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}
