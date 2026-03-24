package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/aprimr/chautari/utils"
	"github.com/aprimr/chautari/validation"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Get auth header and trim prefix
		authHeader := r.Header.Get("Authorization")
		if validation.IsEmptyString(authHeader) {
			utils.SendError(w, "Auth header missing", http.StatusUnauthorized)
			return
		}

		jwtToken := strings.TrimPrefix(authHeader, "Bearer ")
		if validation.IsEmptyString(jwtToken) {
			utils.SendError(w, "Jwt token missing", http.StatusUnauthorized)
			return
		}

		// Verify JWT
		jwtClaims, err := utils.VerifyToken(jwtToken)
		if err != nil {
			utils.SendError(w, "JWT token tampered", http.StatusUnauthorized)
			return
		}

		// Send jwt with next request
		ctx := context.WithValue(r.Context(), "uid", jwtClaims.Uid)
		ctx = context.WithValue(ctx, "username", jwtClaims.Username)
		ctx = context.WithValue(ctx, "email", jwtClaims.Email)

		// call next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
