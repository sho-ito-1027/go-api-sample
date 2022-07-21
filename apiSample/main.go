package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.RequestURI == "/favicon.ico" {
		return
	}
	_, err := fmt.Fprintf(w, "Welcome to the HomePage")
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
	fmt.Println("Endpoint Hit: homePage")
}

func articles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getArticles(w, r)
	case "POST":
		postArticles(w, r)
	default:
		w.WriteHeader(405)
	}
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	articles := make([]Article, 0)
	for i := 0; i < 10; i++ {
		title := "article %d"
		articles = append(articles, Article{Title: fmt.Sprintf(title, i), Desc: "description", Content: "content"})
	}
	fmt.Println("Endpoint Hit: articles")
	err := json.NewEncoder(w).Encode(articles)
	if err != nil {
		log.Fatalf("%v", err)
	}
	w.WriteHeader(200)
}

func postArticles(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	fmt.Println(uri)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	var article Article
	err = json.Unmarshal(body, &article)
	if err != nil {
		log.Fatalf("%v", err)
	}
	fmt.Println(article.Title)
	fmt.Println(article.Desc)
	fmt.Println(article.Content)
	w.WriteHeader(201)
}

func handleRequests() {
	//server := http.Server{
	//	Addr: ":8080",
	//}
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", articles)
	//_ = server.ListenAndServe()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("api")
	handleRequests()
}
