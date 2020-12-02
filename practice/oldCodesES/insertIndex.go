package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	// Import the Elasticsearch library packages
	elasticsearch "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// Declare a struct for Elasticsearch fields
type ElasticDocs struct {
	Name      string
	Id        int
	Available bool
	NewField  string
}

func main() {

	// Allow for custom formatting of log output
	log.SetFlags(0)

	// Create a context object for the API calls
	ctx := context.Background()

	// enabling Insecure Skip Verify for https first assign
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

	// Have the client instance return a response
	// res, err := client.Info()
	_, err = client.Info()

	// Deserialize the response into a map.
	if err != nil {
		log.Fatalf("client.Info() ERROR:", err)
	} //else {
	// 	log.Printf("client response:", res)
	// }

	// Declare documents to be indexed using struct
	doc1 := ElasticDocs{}
	doc1.Name = "name3"
	doc1.Id = 1233
	doc1.Available = true
	doc1.NewField = "addingnewfield3"

	// Marshal Elasticsearch document struct objects to JSON string
	docStr1 := jsonStruct(doc1)

	// Instantiate a request object
	// it can update or add new document, related to id that i added in Itoa
	req := esapi.IndexRequest{
		Index:      "some_index",     // Index name
		DocumentID: strconv.Itoa(11), // should add the id value to insert new, otherwise rewrite last ite
		Body:       strings.NewReader(docStr1),
		Refresh:    "true",
	}
	fmt.Println(reflect.TypeOf(req))

	// Return an API response object from request
	res, err := req.Do(ctx, client)
	if err != nil {
		log.Fatalf("IndexRequest ERROR: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("%s ERROR indexing document ID=%d", res.Status(), 1)
	} else {

		// Deserialize the response into a map.
		var resMap map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&resMap); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			log.Printf("\nIndexRequest() RESPONSE:")
			// Print the response status and indexed document version.
			fmt.Println("Status:", res.Status())
			fmt.Println("Result:", resMap["result"])
			fmt.Println("\n")
		}
	}
}

// A function for marshaling structs to JSON string
func jsonStruct(doc ElasticDocs) string {

	// Create struct instance of the Elasticsearch fields struct object
	docStruct := &ElasticDocs{
		Name:      doc.Name,
		Id:        doc.Id,
		Available: doc.Available,
		NewField:  doc.NewField,
	}
	// Marshal the struct to JSON and check for errors
	b, err := json.Marshal(docStruct)
	if err != nil {
		fmt.Println("json.Marshal ERROR:", err)
		return string(err.Error())
	}
	return string(b)
}
