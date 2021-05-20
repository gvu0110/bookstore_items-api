package services

import (
	"github.com/gvu0110/bookstore_items-api/domain/items"
	"github.com/gvu0110/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemServiceInterface = &itemService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RESTError)
	Get(string) (*items.Item, *rest_errors.RESTError)
}

type itemService struct{}

func (s *itemService) Create(items.Item) (*items.Item, *rest_errors.RESTError) {
	return nil, nil
}

func (s *itemService) Get(string) (*items.Item, *rest_errors.RESTError) {
	return nil, nil
}
