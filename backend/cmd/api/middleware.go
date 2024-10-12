package main

import (
	"fmt"
	"net/http"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Origin")
		w.Header().Set("Vary", "Access-Control-Request-Method")
		w.Header().Set("Vary", "Access-Control-Request-Headers")

		origin := r.Header.Get("Origin")

		if origin != "" && app.config.origin == origin {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			app.logger.PrintInfo(fmt.Sprintf("cors from: %s", origin), nil)

			if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
				w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
				w.Header().Set("Access-Control-Max-Age", "60")
			}
		} else {
			app.logger.PrintInfo(fmt.Sprintf("cors failed, origin: %s", origin), nil)
		}

		next.ServeHTTP(w, r)
	})
}
