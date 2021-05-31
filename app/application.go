package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gvu0110/bookstore_items-api/clients/elasticsearch"
	"github.com/gvu0110/bookstore_utils-go/logger"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	mapURLs()
	logger.Info("Starting the application ...")
	srv := &http.Server{
		Handler: router,
		Addr:    "localhost:8082",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
