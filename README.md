# bookstore_items-api

1. Start up ElasticSearch
```shell
docker-compose up -d es01 es02 es03
```
2. Run locally item-api app
```shell
export ES_ADDRESS=localhost:9200
go run main.go
```