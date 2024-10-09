package main

import (
	"context"
	"github.com/Ddarli/gym/common"
	"github.com/Ddarli/gym/common/logger"
	"github.com/Ddarli/gym/common/tracer"
	"github.com/Ddarli/gym/kafka"
	"github.com/Ddarli/gym/shceduleservice/db"
	"github.com/Ddarli/gym/shceduleservice/models"
	"github.com/Ddarli/gym/shceduleservice/repository"
	"github.com/Ddarli/gym/shceduleservice/services"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"net"
	"os"
)

var log = logger.GetLogger()

func newService() (models.ScheduleServiceServer, *services.KafkaService) {
	conf := common.LoadConfig()
	db, err := db.NewPostgresConnection(conf)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewPostgresScheduleRepository(db)
	classService := services.NewScheduleService(repo)
	kafkaService := services.NewKafkaService(repo)
	return classService, kafkaService
}

func startServer(service models.ScheduleServiceServer) {
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
	models.RegisterScheduleServiceServer(grpcServer, service)
	log.Infof("Serving gRPC on %s", address)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}

func main() {
	cleanup := tracer.InitTracer("schedule-service")
	defer cleanup()
	service, kafkaService := newService()
	go startServer(service)
	ctx := context.Background()
	brokers := []string{"localhost:9095", "localhost:9096", "localhost:9097"}
	topic := "class-checked-events"
	group := "schedule-service-group"

	err := kafka.StartConsuming(ctx, brokers, topic, group, kafkaService.ProcessAvailabilityCheck, logger.GetLogger())
	if err != nil {
		log.Fatal(err)
	}
	<-ctx.Done()
}
