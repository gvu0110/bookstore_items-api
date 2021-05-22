package controllers

import (
	"net/http"

	"github.com/gvu0110/bookstore_items-api/domain/items"
	"github.com/gvu0110/bookstore_items-api/services"
	"github.com/gvu0110/bookstore_items-api/utils/http_utils"
	"github.com/gvu0110/bookstore_oauth-go/oauth"
)

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
	if err := oauth.AuthenticateRequest(r); err != nil {
		http_utils.ResponseRESTError(w, *err)
		return
	}

	// TODO: Unmarshall request into the item struct
	item := items.Item{
		Seller: oauth.GetCallerID(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		http_utils.ResponseRESTError(w, *err)
		return
	}

	http_utils.ResponseJSON(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
}
