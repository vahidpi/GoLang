package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"reflect"

	// Import the Elasticsearch library packages
	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

// Declare a struct for Elasticsearch fields
// this is for test, the real structure should be open to get new fields
type ESDocs struct {
	Name      string
	Id        int
	Available bool
}

func main() {
	// Allow for custom formatting of log output
	log.SetFlags(0)

	// Create a context object for the API calls
	ctx := context.Background()

	// Create a mapping for the Elasticsearch documents
	var (
		docMap map[string]interface{}
	)
	fmt.Println("docMap:", docMap)
	fmt.Println("docMap TYPE:", reflect.TypeOf(docMap))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Declare an Elasticsearch configuration
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://127.0.0.1:9200",
		},
		// Username: "user",
		// Password: "pass",
		Transport: tr,
	}

	// Instantiate a new Elasticsearch client object instance
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println("Elasticsearch connection error:", err)
	}

	// Declare empty array for the document strings
	var docs []string

	// Declare documents to be indexed using struct
	doc1 := ElasticDocs{}
	doc1.Name = "Esm"
	doc1.Id = 123456
	doc1.Available = true
}
