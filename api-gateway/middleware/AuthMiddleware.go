package handlers

import (
	"context"
	"github.com/Ddarli/gym/userservice/models"
	"net/http"
	"strings"
)

func TokenAuthMiddleware(userService models.UserServiceClient) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if len(token) == 0 || !strings.HasPrefix(token, "Bearer ") {
				http.Error(w, "Missing or malformed token", http.StatusUnauthorized)
				return
			}
			token = strings.TrimPrefix(token, "Bearer ")
			ctx := context.Background()
			response, err := userService.VerifyToken(ctx, &models.VerifyTokenRequest{Token: token})
			if err != nil || !response.GetResult() {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
