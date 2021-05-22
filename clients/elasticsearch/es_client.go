package elasticsearch

import (
	"context"
	"time"

	"github.com/olivere/elastic"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	setClient(client *elastic.Client)
	Index(interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("localhost:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	Client.setClient(client)
}

func (c *esClient) Index(interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	return c.client.Index().Do(ctx)
}
