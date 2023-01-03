package bigquery_mock

import (
	"context"
	"fmt"
	"github.com/goccy/bigquery-emulator/server"
	"github.com/goccy/bigquery-emulator/types"
	"google.golang.org/api/iterator"
	"path/filepath"
	"testing"
)

func TestMockBigQueryWithYamlFixtureLoad(t *testing.T) {
	projectName := "test"
	// because we have two records in the yaml file
	expectedResultNo := 2
	ctx := context.Background()
	bqClient, err := MockBigQuery(projectName, server.YAMLSource(filepath.Join("fixtures", "sample_bigquery_fixture.yaml")))
	if err != nil {
		t.Fatalf("error to mock bigquery %v", err)
	}
	query := bqClient.Query("select * from test.dataset_1.table_1")
	job, err := query.Run(ctx)
	if err != nil {
		t.Fatalf("error to run query %v", err)
	}
	it, err := job.Read(ctx)
	if err != nil {
		t.Fatalf("error to run query %v", err)
	}
	type result struct {
	}
	var res []result
	for {
		var row result
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			t.Fatalf("bigquery iterator next failed %v", err)
		}
		res = append(res, row)
	}
	if len(res) != expectedResultNo {
		t.Errorf("result mismatching epected %d, got %d", expectedResultNo, len(res))
	}
}
func TestMockBigQueryWithInlineSourcing(t *testing.T) {
	expectedResultNo := 2
	ctx := context.Background()

	projectName := "test2"
	datasetName := "dataset_1"
	tableName := "table_1"
	source := server.StructSource(
		types.NewProject(
			projectName,
			types.NewDataset(
				datasetName,
				types.NewTable(
					tableName,
					[]*types.Column{
						types.NewColumn("id", types.INTEGER),
						types.NewColumn("field_a", types.BOOL),
						types.NewColumn("field_b", types.BOOL),
						types.NewColumn("field_c", types.STRING),
						types.NewColumn("field_e", types.STRING),
					},
					types.Data{
						{
							"id":      1,
							"field_a": false,
							"field_b": false,
							"field_c": "ok",
							"field_d": "yuseferi",
							"field_e": "A",
						}, {
							"id":      3,
							"field_a": false,
							"field_b": true,
							"field_c": "nok",
							"field_d": "golang",
							"field_e": "A",
						},
					},
				),
			),
		),
	)
	bqClient, err := MockBigQuery(projectName, source)
	if err != nil {
		t.Fatalf("error to mock bigquery %v", err)
	}
	query := bqClient.Query(fmt.Sprintf("select * from %s.%s.%s", projectName, datasetName, tableName))
	job, err := query.Run(ctx)
	if err != nil {
		t.Fatalf("error to run query %v", err)
	}
	it, err := job.Read(ctx)
	if err != nil {
		t.Fatalf("error to run query %v", err)
	}
	type result struct {
	}
	var res []result
	for {
		var row result
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			t.Fatalf("bigquery iterator next failed %v", err)
		}
		res = append(res, row)
	}
	if len(res) != expectedResultNo {
		t.Errorf("result mismatching epected %d, got %d", expectedResultNo, len(res))
	}
}
