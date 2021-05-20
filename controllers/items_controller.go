package controllers

import "net/http"

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Search(http.ResponseWriter, *http.Request)
}

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
}
