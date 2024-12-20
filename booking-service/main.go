package main

import (
	"context"
	"github.com/Ddarli/gym/bookingservice/db"
	"github.com/Ddarli/gym/bookingservice/models"
	"github.com/Ddarli/gym/bookingservice/repository"
	"github.com/Ddarli/gym/bookingservice/services"
	"github.com/Ddarli/gym/kafka"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"net"
	"os"
)

var log = logger.GetLogger()

func main() {
	cleanup := tracer.InitTracer("booking-service")
	defer cleanup()
	service, kafkaService := newService()
	go startServer(service)
	ctx := context.Background()
	brokers := []string{"localhost:9095", "localhost:9096", "localhost:9097"}
	topic := "schedule-checked-events"
	group := "booking-service-group"

	err := kafka.StartConsuming(ctx, brokers, topic, group, kafkaService.ProcessAvailabilityCheck, logger.GetLogger())
	if err != nil {
		log.Fatal(err)
	}
	<-ctx.Done()
}

func newService() (models.BookingServiceServer, *services.KafkaService) {
	conf := common.LoadConfig()
	db, err := db.NewPostgresConnection(conf)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewPostgresBookingRepository(db)
	classService := services.NewBookingService(repo)
	kafkaService := services.NewKafkaService(repo)
	return classService, kafkaService
}

func startServer(service models.BookingServiceServer) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	address := os.Getenv("ADDRESS")
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	models.RegisterBookingServiceServer(grpcServer, service)
	log.Infof("Serving gRPC on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}
