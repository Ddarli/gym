package handlers

import (
	"context"
	"encoding/json"
	bookingmodel "github.com/Ddarli/gym/bookingservice/models"
	handlers "github.com/Ddarli/gym/gateway/middleware"
	"github.com/Ddarli/gym/userservice/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

type handler struct {
	userService    models.UserServiceClient
	bookingService bookingmodel.BookingServiceClient
}

func NewHandler(userServiceClient models.UserServiceClient, bookingService bookingmodel.BookingServiceClient) *handler {
	return &handler{
		userService:    userServiceClient,
		bookingService: bookingService,
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
		bookingModel := bookingmodel.ToBookingModel(booking)
		if err := json.NewEncoder(w).Encode(bookingModel); err != nil {
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
