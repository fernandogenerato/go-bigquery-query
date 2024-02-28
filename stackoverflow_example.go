package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

func main() {

	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}
	defer client.Close()

	rows, err := q(ctx, client)
	if err != nil {
		log.Fatal(err)
	}
	if err := printResults(os.Stdout, rows); err != nil {
		log.Fatal(err)
	}
}

// query returns a row iterator suitable for reading query results.
func q(ctx context.Context, client *bigquery.Client) (*bigquery.RowIterator, error) {

	query := client.Query(
		`SELECT
                        CONCAT(
                                'https://stackoverflow.com/questions/',
                                CAST(id as STRING)) as url,
                        view_count
                FROM ` + "`bigquery-public-data.stackoverflow.posts_questions`" + `
                WHERE tags like '%google-bigquery%'
                ORDER BY view_count DESC
                LIMIT 10;`)
	return query.Read(ctx)
}

type StackOverflowRow struct {
	URL       string `bigquery:"url"`
	ViewCount int64  `bigquery:"view_count"`
}

// printResults prints results from a query to the Stack Overflow public dataset.
func printResults(w io.Writer, iter *bigquery.RowIterator) error {
	for {
		var row StackOverflowRow
		err := iter.Next(&row)
		if err == iterator.Done {
			return nil
		}
		if err != nil {
			return fmt.Errorf("error iterating through results: %v", err)
		}

		fmt.Fprintf(w, "url: %s views: %d\n", row.URL, row.ViewCount)
	}
}
