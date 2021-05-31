package services

import (
	"github.com/gvu0110/bookstore_items-api/domain/items"
	"github.com/gvu0110/bookstore_items-api/domain/queries"
	"github.com/gvu0110/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemServiceInterface = &itemService{}
)

type itemServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RESTError)
	Get(string) (*items.Item, rest_errors.RESTError)
	Search(queries.ESQuery) ([]items.Item, rest_errors.RESTError)
	Delete(string) rest_errors.RESTError
}

type itemService struct{}

func (s *itemService) Create(item items.Item) (*items.Item, rest_errors.RESTError) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Get(id string) (*items.Item, rest_errors.RESTError) {
	item := items.Item{ID: id}
	if err := item.Get(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemService) Search(query queries.ESQuery) ([]items.Item, rest_errors.RESTError) {
	dao := items.Item{}
	return dao.Search(query)
}

func (s *itemService) Delete(id string) rest_errors.RESTError {
	item := items.Item{ID: id}
	if err := item.Delete(); err != nil {
		return err
	}
	return nil
}
