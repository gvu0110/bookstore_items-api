package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapURLs()

	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:8082",
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
