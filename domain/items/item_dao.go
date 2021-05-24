package items

import (
	"errors"

	"github.com/gvu0110/bookstore_items-api/clients/elasticsearch"
	"github.com/gvu0110/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() *rest_errors.RESTError {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerRESTError("Error when trying to save item", errors.New("database error"))
	}
	i.ID = result.Id
	return nil
}
