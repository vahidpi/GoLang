package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type TestMap struct {
	Title string
	Text  string
}

type WebAggPage struct {
	Title string
	Text  map[int]string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>TestPage</h1>")
}

func webPageHandler(w http.ResponseWriter, r *http.Request) {

	//test_map := make(map[string]TestMap)

	map_2 := map[int]string{

		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	p := WebAggPage{Title: "Test", Text: map_2}
	t, _ := template.ParseFiles("testPage.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/web/", webPageHandler)
	http.ListenAndServe(":8001", nil)
}
