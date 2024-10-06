package main

import (
	"github.com/Ddarli/gym/classservice/db"
	"github.com/Ddarli/gym/classservice/models"
	"github.com/Ddarli/gym/classservice/repository"
	"github.com/Ddarli/gym/classservice/services"
	"github.com/Ddarli/gym/common"
	"github.com/Ddarli/gym/common/logger"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"net"
	"os"
)

var log = logger.GetLogger()

func newService() models.ClassServiceServer {
	conf := common.LoadConfig()
	db, err := db.NewPostgresConnection(conf)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewPostgresClassRepository(db)
	classService := services.NewClassService(repo)
	return classService
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
	models.RegisterClassServiceServer(grpcServer, service)
	log.Infof("Serving gRPC on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}
