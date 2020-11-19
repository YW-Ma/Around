package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

const (
	ES_URL = "http://10.170.0.2:9200"
)

func readFromES(query elastic.Query, index string) (*elastic.SearchResult, error) {
	// Step 1. connection
	fmt.Println("step into readFromES")
	client, err := elastic.NewClient(
		elastic.SetURL(ES_URL),
		elastic.SetBasicAuth("elastic", "MatrixMayaowei123"))
	if err != nil {
		fmt.Println("client creation error")
		return nil, err
	}
	// Step 2. Search using arguments
	// searchResult是搜索结果, search with a term query
	searchResult, err := client.Search().
		Index(index).            // search in index ‘index’
		Query(query).            // specify the query
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		return nil, err
	}
	// Step 3. Return the result
	return searchResult, nil
}
