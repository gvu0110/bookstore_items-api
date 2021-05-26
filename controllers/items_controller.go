package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gvu0110/bookstore_items-api/domain/items"
	"github.com/gvu0110/bookstore_items-api/domain/queries"
	"github.com/gvu0110/bookstore_items-api/services"
	"github.com/gvu0110/bookstore_items-api/utils/http_utils"
	"github.com/gvu0110/bookstore_oauth-go/oauth"
	"github.com/gvu0110/bookstore_utils-go/rest_errors"
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

	sellerID := oauth.GetCallerID(r)
	if sellerID == 0 {
		restErr := rest_errors.NewUnauthorizedRESTError("Unable to retrieve information from the given access token")
		http_utils.ResponseRESTError(w, *restErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restErr := rest_errors.NewBadRequestRESTError("Invalid request body")
		http_utils.ResponseRESTError(w, *restErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		restErr := rest_errors.NewBadRequestRESTError("Invalid item JSON body")
		http_utils.ResponseRESTError(w, *restErr)
		return
	}

	itemRequest.Seller = sellerID

	result, createErr := services.ItemsService.Create(itemRequest)
	if err != nil {
		http_utils.ResponseRESTError(w, *createErr)
		return
	}

	http_utils.ResponseJSON(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := strings.TrimSpace(vars["id"])

	item, err := services.ItemsService.Get(itemID)
	if err != nil {
		http_utils.ResponseRESTError(w, *err)
		return
	}
	http_utils.ResponseJSON(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		restErr := rest_errors.NewBadRequestRESTError("Invalid JSON body")
		http_utils.ResponseRESTError(w, *restErr)
		return
	}
	defer r.Body.Close()

	var query queries.ESQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		restErr := rest_errors.NewBadRequestRESTError("Invalid JSON body")
		http_utils.ResponseRESTError(w, *restErr)
		return
	}

	items, searchErr := services.ItemsService.Search(query)
	if searchErr != nil {
		http_utils.ResponseRESTError(w, *searchErr)
		return
	}
	http_utils.ResponseJSON(w, http.StatusOK, items)
}
