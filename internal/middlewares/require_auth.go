package middlewares

import (
	"context"
	"errors"
	"github.com/muktiarafi/myriadcode-backend/internal/apierror"
	"github.com/muktiarafi/myriadcode-backend/internal/helpers"
	"net/http"
)

func RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			helpers.SendError(w, apierror.NewBadRequestError(err, "Missing cookie"))
			return
		}

		token, payload, err := helpers.ParseToken(cookie.Value)
		if err != nil {
			helpers.SendError(w, err)
			return
		}

		if !token.Valid {
			helpers.SendError(w, apierror.NewBadRequestError(
				errors.New("invalid token"), "invalid token"))
		}

		ctx := context.WithValue(r.Context(), "user", payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
