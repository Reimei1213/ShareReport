package auth

import (
	"context"
	"net/http"
)

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("")

			if err != nil || c == nil {
				next.ServeHTTP(w, r)
				return
			}

			userId := "1"
			// userId, err := validateAndGetUserID(c)
			// if err != nil {
			// 	http.Error(w, "Invalid cookie", http.StatusForbidden)
			// }
			ctx := context.WithValue(r.Context(), "userId", userId)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
