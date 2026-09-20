package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const IDKey = "id"

func ParseID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil || id <= 0 {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), IDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetID(ctx context.Context) int {
	return ctx.Value(IDKey).(int)
}
