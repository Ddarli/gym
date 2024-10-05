package handlers

import (
	"context"
	"encoding/json"
	bookingmodel "github.com/Ddarli/gym/bookingservice/models"
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

func NewHandler(usesrService models.UserServiceClient, bookingService bookingmodel.BookingServiceClient) *handler {
	return &handler{
		userService:    usesrService,
		bookingService: bookingService,
	}
}

func (h *handler) RegisterRoutes(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(5 * time.Second))
	h.registerUserHandlers(r)
	h.registerBookingHandlers(r)
}

func (h *handler) registerUserHandlers(r *chi.Mux) {
	r.Get("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		id := chi.URLParam(r, "userId")
		req := models.GetUserRequest{
			Id: id,
		}
		user, err := h.userService.GetUser(ctx, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
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
	})
	r.Patch("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		id := chi.URLParam(r, "userId")
		user := models.User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		user.Id = id
		req := models.UpdateUserRequest{Id: id, User: &user}
		savedUser, err := h.userService.UpdateUser(ctx, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(savedUser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	r.Delete("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		id := chi.URLParam(r, "userId")
		req := models.DeleteUserRequest{Id: id}
		res, err := h.userService.DeleteUser(ctx, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func (h *handler) registerBookingHandlers(r *chi.Mux) {
	r.Get("/bookings/{bookingId}", func(w http.ResponseWriter, r *http.Request) {
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
	})
	r.Post("/bookings", func(w http.ResponseWriter, r *http.Request) {
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
	})
	r.Delete("/bookings/{booingId}", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		id := chi.URLParam(r, "bookingId")
		req := bookingmodel.CancelBookingRequest{Id: id}
		resp, err := h.bookingService.CancelBooking(ctx, &req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
