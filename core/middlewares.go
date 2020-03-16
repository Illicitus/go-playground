package core

import (
	"context"
	"github.com/go-pg/pg/v9"
	"net/http"
)

//func DatabaseMiddleware(db *pg.DB) func(http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			ctx := context.WithValue(r.Context(), "db", db)
//			next.ServeHTTP(w, r.WithContext(ctx))
//		})
//	}
//}

func JwtAuthMiddleware(db *pg.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Add db to context
			ctx := context.WithValue(r.Context(), "db", db)

			// Add Bearer jwt token or set it as nil
			jwtToken := r.Header.Get("Authorization")
			if jwtToken != "" {
				ctx = context.WithValue(ctx, "userEmail", jwtToken)
			} else {
				ctx = context.WithValue(ctx, "userEmail", nil)
			}

			// Add updated context to handler request
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
