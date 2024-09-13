package main

import (
	"net/http"
)

func AuthChecker(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// Controllo della validita del token

		handler.ServeHTTP(writer, request)
	})
}
