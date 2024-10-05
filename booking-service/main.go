package main

import (
	"github.com/Ddarli/gym/bookingservice/db"
	"github.com/Ddarli/gym/bookingservice/models"
	"github.com/Ddarli/gym/bookingservice/repository"
	"github.com/Ddarli/gym/bookingservice/services"
	"github.com/Ddarli/gym/common"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func newService() models.BookingServiceServer {
	conf := common.LoadConfig()
	//logger, _ := zap.NewDevelopment()
	//defer logger.Sync()
	db, err := db.NewPostgresConnection(conf)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewPostgresBookingRepository(db)
	userService := services.NewBookingService(repo)
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
	models.RegisterBookingServiceServer(grpcServer, service)
	log.Printf("Serving gRPC on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}
