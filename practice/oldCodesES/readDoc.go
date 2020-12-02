package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	// Import the Elasticsearch library packages
	"github.com/elastic/go-elasticsearch/v8"
)

func main() {

	// Allow for custom formatting of log output
	log.SetFlags(0)

	// Create a context object for the API calls
	ctx := context.Background()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Declare an Elasticsearch configuration
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		//Username:  "user",
		//Password:  "pass",
		Transport: tr,
	}

	// Instantiate a new Elasticsearch client object instance
	client, err := elasticsearch.NewClient(cfg)

	// Exit the system if connection raises an error
	if err != nil {
		log.Fatalf("Elasticsearch connection error:", err)
	}

	// Instantiate a mapping interface for API response
	var mapResp map[string]interface{}
	var buf bytes.Buffer

	// More example query strings to read and pass to Search()
	//var query = `{"query": {"match_all" : {}},"size": 5}` // size of docs that will read
	var query = `{"query": {"match_all" : {}}}`
	// query = `{"query": {"term" : {"SomeBool": true}},"size": 3}`
	// query = `{"query": {"match" : {"SomeStr": "Value"}},"size": 3}`

	// Concatenate a string from query for reading
	var b strings.Builder
	b.WriteString(query)
	read := strings.NewReader(b.String())

	// Attempt to encode the JSON query and look for errors
	if err := json.NewEncoder(&buf).Encode(read); err != nil {
		log.Fatalf("Error encoding query: %s", err)

		// Query is a valid JSON object
	} else {

		// Pass the JSON query to the Golang client's Search() method
		res, err := client.Search(
			client.Search.WithContext(ctx),
			client.Search.WithIndex("some_index"),
			client.Search.WithBody(read),
			client.Search.WithTrackTotalHits(true),
			client.Search.WithPretty(),
		)

		// Check for any errors returned by API call to Elasticsearch
		if err != nil {
			log.Fatalf("Elasticsearch Search() API ERROR:", err)

			// If no errors are returned, parse esapi.Response object
		} else {
			// Close the result body when the function call is complete
			defer res.Body.Close()

			// Decode the JSON response and using a pointer
			if err := json.NewDecoder(res.Body).Decode(&mapResp); err != nil {
				log.Fatalf("Error parsing the response body: %s", err)

				// If no error, then convert response to a map[string]interface
			} else {
				fmt.Println("Read all docs ____ :", "\n")

				// Iterate the document "hits" returned by API call
				for _, hit := range mapResp["hits"].(map[string]interface{})["hits"].([]interface{}) {

					// Parse the attributes/fields of the document
					doc := hit.(map[string]interface{})

					// The "_source" data is another map interface nested inside of doc
					source := doc["_source"]
					fmt.Println("doc _source:", reflect.TypeOf(source))

					// Get the document's _id and print it out along with _source data
					docID := doc["_id"]
					fmt.Println("docID:", docID)
					fmt.Println("_source:", source, "\n")
				} // end of response iteration

			}
		}
	}
} // end of main() func
