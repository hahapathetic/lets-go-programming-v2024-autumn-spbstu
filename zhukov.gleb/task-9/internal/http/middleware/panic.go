package middleware

import (
	"log"
	"net/http"

	"task-9/internal/http/handler"
)

func PanicHandler(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				handler.WriteJSONServer(w, map[string]string{"message": "Internal server error"}, http.StatusInternalServerError)
				logger.Fatal("recovered panic")
			}
		}()
		next.ServeHTTP(w, r)
	})
}
