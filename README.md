# bigquery-mock
At the moment [google bigquery client](cloud.google.com/go/bigquery) does not provide an interface for it's client and wrapping all requirement in custom interfaces needs lots of codes and energy and keeping it uptodated with all Google Bigquery could be hard. (something which has been done in [this archived](https://github.com/googleapis/google-cloud-go-testing) project).   
This project is a simple functionality to mock Google Bigquery client for testing purposes. technically it's using [bigquery-emulator](https://github.com/goccy/bigquery-emulator) under the hood to provide the basics with minimum effort.

#### Usage
add it to your project

`go get github.com/yuseferi/bigquery-mock`

then import it in your tests

`import bigquery_mock "github.com/yuseferi/bigquery-mock"`

There are two options to define your schemas and add some test data to your mocked bigquery service.
- Yaml file loader
- On the fly sourcing

#### Yaml file loader
it automatically creates datasets, tables. it's schema and fill tables with data.

sample yaml file: 
```yaml
projects:
  - id: test
    datasets:
      - id: dataset_1
        tables:
          - id: table_1
            columns:
              - name: field_a
                type: STRING
              - name: field_b
                type: STRING
              - name: field_c
                type: INT
              - name: field_d
                type: BOOL
            data:
              - field_a: "record_1_a"
                field_b: "record_1_b"
                field_c: 333
                field_d: false
              - field_a: "record_2_a"
                field_b: "record_2_b"
                field_c: 222
                field_d: true


```
then load it by simple mocking
```go
bqClient, err := bigquery_mock.MockBigQuery(projectName, server.YAMLSource(filepath.Join("testdata", "sample_bigquery_fixture.yaml")))
if err != nil {
    t.Fatalf("error to mock bigquery")
}
query := bqClient.Query("select * from test.dataset_1.table_1")
```


#### On the fly sourcing
in this solution you can create datasets, tables and schemas and fill tables on the fly in your tests. it means you don't need to define you dataset, schema in yaml file.

example: 

```go
projectName := "test"
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
                    types.NewColumn("field_f", types.STRING),
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
    t.Fatalf("error to mock bigquery")
}
query := bqClient.Query(fmt.Sprintf("select * from %s.%s.%s", projectName, datasetName, tableName))
```
As you can see in the above sample, you can define your dataset,tables and on your test.


**feel free to contribute with creating PR and make it better :)** 

### License
The GNU General Public License v3.0
