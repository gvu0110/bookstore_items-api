package items

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gvu0110/bookstore_items-api/clients/elasticsearch"
	"github.com/gvu0110/bookstore_items-api/domain/queries"
	"github.com/gvu0110/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
	typeItem   = "_doc"
)

func (i *Item) Save() *rest_errors.RESTError {
	result, err := elasticsearch.Client.Index(indexItems, typeItem, i)
	if err != nil {
		return rest_errors.NewInternalServerRESTError("Error when trying to save item", errors.New("database error"))
	}
	i.ID = result.Id
	return nil
}

func (i *Item) Get() *rest_errors.RESTError {
	itemID := i.ID
	result, err := elasticsearch.Client.Get(indexItems, typeItem, i.ID)
	if err != nil {
		return rest_errors.NewInternalServerRESTError(fmt.Sprintf("Error when trying to get id %s", i.ID), errors.New("database error"))
	}
	if !result.Found {
		return rest_errors.NewNotFoundRESTError(fmt.Sprintf("No item found with id %s", i.ID))
	}

	bytes, err := result.Source.MarshalJSON()
	if err != nil {
		return rest_errors.NewInternalServerRESTError("Error when trying to parse database response", errors.New("database error"))
	}

	if err := json.Unmarshal(bytes, i); err != nil {
		return rest_errors.NewInternalServerRESTError("Error when trying to unmarshal database bytes response", errors.New("database error"))
	}
	i.ID = itemID
	return nil
}

func (i *Item) Search(query queries.ESQuery) ([]Item, *rest_errors.RESTError) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, rest_errors.NewInternalServerRESTError("Error when trying to search documents", errors.New("database error"))
	}

	items := make([]Item, result.TotalHits())
	for i, hit := range result.Hits.Hits {
		bytes, err := hit.Source.MarshalJSON()
		if err != nil {
			return nil, rest_errors.NewInternalServerRESTError("Error when trying to parse database response", errors.New("database error"))
		}
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, rest_errors.NewInternalServerRESTError("Error when trying to unmarshal database bytes response", errors.New("database error"))
		}
		item.ID = hit.Id
		items[i] = item
	}

	if len(items) == 0 {
		return nil, rest_errors.NewNotFoundRESTError("No item found")
	}
	return items, nil
}
