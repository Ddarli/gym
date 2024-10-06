package main

import (
	bookingmodel "github.com/Ddarli/gym/bookingservice/models"
	classservice "github.com/Ddarli/gym/classservice/models"
	"github.com/Ddarli/gym/gateway/config"
	"github.com/Ddarli/gym/gateway/handlers"
	trainerservice "github.com/Ddarli/gym/trainerservice/models"
	"github.com/Ddarli/gym/userservice/models"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
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
	defer classServiceConnection.Close()

	userServiceClient := models.NewUserServiceClient(userServiceConnection)
	bookingServiceClient := bookingmodel.NewBookingServiceClient(bookingServiceConnection)
	classServiceClient := classservice.NewClassServiceClient(classServiceConnection)
	trainerServiceClient := trainerservice.NewTrainerServiceClient(trainerServiceConnection)

	handler := handlers.NewHandler(userServiceClient, bookingServiceClient, classServiceClient, trainerServiceClient)

	r := chi.NewRouter()
	handler.RegisterRoutes(r)
	if err := http.ListenAndServe(httpAdd, r); err != nil {
		log.Fatal(err)
	}
}
