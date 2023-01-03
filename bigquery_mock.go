package bigquery_mock

import (
	"cloud.google.com/go/bigquery"
	"context"
	"github.com/goccy/bigquery-emulator/server"
	"google.golang.org/api/option"
)

// MockBigQuery is a basic function that create a bigquery mocked client for you.
// projectName is project name which you wanted to be considered as bigquery project in the mocked service.
// source is the source data which you want to be added to the mocked service, it could a yaml source file or on the Fly sourcing see https://github.com/yuseferi/bigquery-mock for more information.
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
