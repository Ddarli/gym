package main

import (
	bookingmodel "github.com/Ddarli/gym/bookingservice/models"
	classservice "github.com/Ddarli/gym/classservice/models"
	"github.com/Ddarli/gym/common/tracer"
	"github.com/Ddarli/gym/gateway/config"
	"github.com/Ddarli/gym/gateway/handlers"
	scheduleservice "github.com/Ddarli/gym/shceduleservice/models"
	trainerservice "github.com/Ddarli/gym/trainerservice/models"
	"github.com/Ddarli/gym/userservice/models"
	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	cleanup := tracer.InitTracer("api-gateway")
	defer cleanup()
	httpAdd := config.HttpAddr
	userServiceConnection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer userServiceConnection.Close()
	bookingServiceConnection, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer bookingServiceConnection.Close()
	classServiceConnection, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer classServiceConnection.Close()
	trainerServiceConnection, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer trainerServiceConnection.Close()
	scheduleServiceConnection, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer scheduleServiceConnection.Close()

	userServiceClient := models.NewUserServiceClient(userServiceConnection)
	bookingServiceClient := bookingmodel.NewBookingServiceClient(bookingServiceConnection)
	classServiceClient := classservice.NewClassServiceClient(classServiceConnection)
	trainerServiceClient := trainerservice.NewTrainerServiceClient(trainerServiceConnection)
	scheduleServiceClient := scheduleservice.NewScheduleServiceClient(scheduleServiceConnection)

	handler := handlers.NewHandler(userServiceClient, bookingServiceClient, classServiceClient, trainerServiceClient,
		scheduleServiceClient)

	r := chi.NewRouter()
	handler.RegisterRoutes(r)
	if err := http.ListenAndServe(httpAdd, otelhttp.NewHandler(r, "gateway")); err != nil {
		log.Fatal(err)
	}
}
