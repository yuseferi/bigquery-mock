package bigquery_mock

import (
	"cloud.google.com/go/bigquery"
	"context"
	"github.com/goccy/bigquery-emulator/server"
	"google.golang.org/api/option"
)

func MockBigQuery(projectName string, sources ...server.Source) (client *bigquery.Client, err error) {
	ctx := context.Background()
	bqServer, err := server.New(server.MemoryStorage)
	if err != nil {
		return nil, err
	}
	if err := bqServer.Load(sources...); err != nil {
		return nil, err
	}
	testServer := bqServer.TestServer()
	client, err = bigquery.NewClient(
		ctx,
		projectName,
		option.WithEndpoint(testServer.URL),
		option.WithoutAuthentication(),
	)
	if err != nil {
		return nil, err

	}
	return client, nil
}
