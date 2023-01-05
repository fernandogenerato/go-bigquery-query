package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

type Item struct {
	Id   int    `bigquery:"id"`
	Name string `bigquery:"name"`
	Test string `bigquery:"name"`
}

const (
	projectID = "bigq-poc"
	dataset   = "data"
	table     = "test_data"
)

func query(ctx context.Context, c *bigquery.Client) {
	q := c.Query(`SELECT * FROM` + "`bigq-poc.data.test_data`")

	it, err := q.Read(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(values)
	}
}
func insert(ctx context.Context, c *bigquery.Client) {
	u := c.Dataset(dataset).Table(table).Inserter()
	i := &Item{
		Id:   2,
		Name: "test3",
		Test: "test",
	}

	err := u.Put(ctx, i)
	if err != nil {
		log.Fatal(err)
	}

}
func main() {

	ctx := context.Background()
	client, _ := bigquery.NewClient(ctx, projectID)
	defer client.Close()

	insert(ctx, client)
	query(ctx, client)

}
