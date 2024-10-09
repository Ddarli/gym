package handlers

import (
	"context"
	"encoding/json"
	bookingmodel "github.com/Ddarli/gym/bookingservice/models"
	classservice "github.com/Ddarli/gym/classservice/models"
	handlers "github.com/Ddarli/gym/gateway/middleware"
	scheduleservice "github.com/Ddarli/gym/shceduleservice/models"
	trainerservice "github.com/Ddarli/gym/trainerservice/models"
	"github.com/Ddarli/gym/userservice/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

type handler struct {
	userService     models.UserServiceClient
	bookingService  bookingmodel.BookingServiceClient
	classService    classservice.ClassServiceClient
	trainerService  trainerservice.TrainerServiceClient
	scheduleService scheduleservice.ScheduleServiceClient
}

func NewHandler(userServiceClient models.UserServiceClient, bookingService bookingmodel.BookingServiceClient,
	classServiceClient classservice.ClassServiceClient, trainerService trainerservice.TrainerServiceClient,
	scheduleService scheduleservice.ScheduleServiceClient) *handler {
	return &handler{
		userService:     userServiceClient,
		bookingService:  bookingService,
		classService:    classServiceClient,
		trainerService:  trainerService,
		scheduleService: scheduleService,
	}
}

func (h *handler) RegisterRoutes(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Group(func(r chi.Router) {
		r.Post("/api/v1/login", h.LoginHandler())
		r.Post("/api/v1/register", h.RegisterHandler())
	})
	r.Group(func(r chi.Router) {
		r.Use(handlers.TokenAuthMiddleware(h.userService))

		r.Get("/api/v1/bookings/{bookingId}", h.GetBookingHandler())
		r.Post("/api/v1/bookings", h.CreateBookingHandler())

		r.Get("/api/v1/schedules/{scheduleId}", h.GetScheduleHandler())
		r.Post("/api/v1/schedules", h.CreateScheduleHandler())

		r.Get("/api/v1/trainers/{trainerId}", h.GetTrainerHandler())

		r.Get("/api/v1/classes/{classId}", h.GetClassHandler())
		r.Get("/api/v1/classes", h.GetClassesHandler())
		r.Post("/api/v1/classes", h.CreateClassHandler())
	})
}

func (h *handler) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		authRequest := models.AuthenticateRequest{}
		if err := json.NewDecoder(r.Body).Decode(&authRequest); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		response, err := h.userService.Authenticate(ctx, &authRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if response.Error == "" && response.Token != "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response.Token))
		}
	}
}
func (h *handler) RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		user := models.User{}

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		request := models.CreateUserRequest{
			Username:    user.Username,
			Password:    user.Password,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
		}
		savedUser, err := h.userService.CreateUser(ctx, &request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if err := json.NewEncoder(w).Encode(savedUser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *handler) GetBookingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		id := chi.URLParam(r, "bookingId")
		request := bookingmodel.GetBookingRequest{Id: id}
		booking, err := h.bookingService.GetBooking(ctx, &request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		//w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(booking); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *handler) CreateBookingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		req := bookingmodel.CreateBookingRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		booking, err := h.bookingService.CreateBooking(ctx, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(booking); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *handler) GetClassHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		classId := chi.URLParam(r, "classId")
		class, err := h.classService.GetClass(ctx, &classservice.GetClassRequest{Id: classId})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(class); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *handler) CreateClassHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		class := &classservice.ClassModel{}
		if err := json.NewDecoder(r.Body).Decode(class); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		request := classservice.CreateClassRequest{
			Name:        class.Name,
			Description: class.Description,
			Capacity:    int32(class.Capacity),
		}
		response, err := h.classService.CreateClass(ctx, &request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		err = json.NewEncoder(w).Encode(&response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *handler) GetClassesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		classes, err := h.classService.GetClasses(ctx, &classservice.GetClassesRequest{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(classes); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *handler) GetTrainerHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		id := chi.URLParam(r, "trainerId")
		trainer, err := h.trainerService.GetTrainer(ctx, &trainerservice.GetTrainerRequest{Id: id})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(trainer); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}

func (h *handler) GetScheduleHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		id := chi.URLParam(r, "scheduleId")
		schedule, err := h.scheduleService.GetSchedule(ctx, &scheduleservice.GetScheduleRequest{Id: id})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(schedule); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *handler) CreateScheduleHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		schedule := &scheduleservice.Schedule{}
		if err := json.NewDecoder(r.Body).Decode(schedule); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		resp, err := h.scheduleService.CreateSchedule(ctx, &scheduleservice.CreateScheduleRequest{
			ClassId:   schedule.ClassId,
			TrainerId: schedule.TrainerId,
			StartTime: schedule.StartTime,
			EndTime:   schedule.EndTime,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}
