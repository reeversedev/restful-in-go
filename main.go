package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Article type
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles type
type Articles []Article

// AllArticles Handler
func AllArticles(w http.ResponseWriter, r *http.Request) {

	articles := Articles{
		Article{Title: "Hello World", Desc: "Hello to the world", Content: "I said, Hello!"},
	}

	fmt.Println("Endpoint Hit : All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

// HomePage handler
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}

func handleRequests() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/articles", AllArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
