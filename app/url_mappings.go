package app

import (
	"net/http"

	"github.com/gvu0110/bookstore_items-api/controllers"
)

func mapURLs() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
	router.HandleFunc("/items/search", controllers.ItemsController.Search).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Delete).Methods(http.MethodDelete)
}
