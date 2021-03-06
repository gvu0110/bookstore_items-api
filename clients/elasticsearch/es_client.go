package elasticsearch

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gvu0110/bookstore_utils-go/logger"
	"github.com/olivere/elastic"
)

const (
	es_env_var = "ES_ADDRESS"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(client *elastic.Client)
	Index(string, string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string, string) (*elastic.GetResult, error)
	Search(string, elastic.Query) (*elastic.SearchResult, error)
	Delete(string, string, string) error
}

type esClient struct {
	client *elastic.Client
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func Init() {
	log := logger.GetLogger()
	client, err := elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("http://%s", os.Getenv(es_env_var))),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetSniff(false),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (c *esClient) Index(index string, docType string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	result, err := c.client.Index().
		Index(index).
		Type(docType).
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to index document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Get(index string, docType string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()
	result, err := c.client.Get().
		Index(index).
		Type(docType).
		Id(id).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to get id %s", id), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := c.client.Search().
		Index(index).
		Query(query).
		RestTotalHitsAsInt(true).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to search document in index %s", index), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Delete(index string, docType string, id string) error {
	ctx := context.Background()
	_, err := c.client.Delete().
		Index(index).
		Type(docType).
		Id(id).
		Refresh("true").
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to delete id %s", id), err)
		return err
	}
	return nil
}
