package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// MODEL - MArticle
type MArticle struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// MODEL - MRequestResponse
type MRequestResponse struct {
	Success bool   `json:"success"`
	Status  string `json:"status"`
}

var DArticles []MArticle

// UTIL
func createNewId() int {
	maxId := 1
	for _, article := range DArticles {
		maxId = article.Id
	}

	return maxId + 1
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage")
	fmt.Println("Endpoint Hit: homePage")
}

// CONTROLLER - Article Controller
func articlesAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: articlesAll")
	json.NewEncoder(w).Encode(DArticles)
}

func articleSingle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println("Endpoint Hit: articleSingle with id " + id)
	for _, article := range DArticles {
		if fmt.Sprint(article.Id) == id {
			json.NewEncoder(w).Encode(article)
			return
		}
	}

	json.NewEncoder(w).Encode(MRequestResponse{
		Success: false,
		Status:  "Article Not Found",
	})
}

func articleCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: articlesCreate")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newArticle MArticle
	json.Unmarshal(reqBody, &newArticle)

	if newArticle.Title == "" {
		json.NewEncoder(w).Encode(MRequestResponse{
			Success: false,
			Status:  "Article Title is Required",
		})
		return
	}
	if newArticle.Desc == "" {
		json.NewEncoder(w).Encode(MRequestResponse{
			Success: false,
			Status:  "Article Desc is Required",
		})
		return
	}
	if newArticle.Content == "" {
		json.NewEncoder(w).Encode(MRequestResponse{
			Success: false,
			Status:  "Article Content is Required",
		})
		return
	}
	newArticle.Id = createNewId()

	DArticles = append(DArticles, newArticle)
	json.NewEncoder(w).Encode(MRequestResponse{
		Success: true,
		Status:  "Article Created Successfully",
	})
}

func articleUpdate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reqBody, _ := ioutil.ReadAll(r.Body)
	id := vars["id"]

	var updateArticle MArticle
	json.Unmarshal(reqBody, &updateArticle)

	if updateArticle.Title == "" {
		json.NewEncoder(w).Encode(MRequestResponse{
			Success: false,
			Status:  "Article Title is Required",
		})
		return
	}
	if updateArticle.Desc == "" {
		json.NewEncoder(w).Encode(MRequestResponse{
			Success: false,
			Status:  "Article Desc is Required",
		})
		return
	}
	if updateArticle.Content == "" {
		json.NewEncoder(w).Encode(MRequestResponse{
			Success: false,
			Status:  "Article Content is Required",
		})
		return
	}

	for index, article := range DArticles {
		if fmt.Sprint(article.Id) == id {
			updateArticle.Id = article.Id
			DArticles = append(DArticles[:index], updateArticle)

			json.NewEncoder(w).Encode(MRequestResponse{
				Success: true,
				Status:  "Article Updated Successfully",
			})
			return
		}
	}

	json.NewEncoder(w).Encode(MRequestResponse{
		Success: false,
		Status:  "Article Not Found",
	})
}

func articleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	fmt.Println("Endpoint Hit: articlesDelete with id " + id)
	for index, article := range DArticles {
		if fmt.Sprint(article.Id) == id {
			DArticles = append(DArticles[:index], DArticles[:index+1]...)
			json.NewEncoder(w).Encode(MRequestResponse{
				Success: true,
				Status:  "Article Deleted Successfully",
			})
			return
		}
	}

	json.NewEncoder(w).Encode(MRequestResponse{
		Success: false,
		Status:  "Article Not Found",
	})
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	// article endpoints
	router.HandleFunc("/articles", articlesAll).Methods("GET")
	router.HandleFunc("/article", articleCreate).Methods("POST")
	router.HandleFunc("/article/{id}", articleSingle).Methods("GET")
	router.HandleFunc("/article/{id}", articleUpdate).Methods("PUT")
	router.HandleFunc("/article/{id}", articleDelete).Methods("DELETE")
	//
	log.Fatal(http.ListenAndServe(":10000", router))
}

func parseData() {
	DArticles = []MArticle{
		{
			Id:      1,
			Title:   "Article One",
			Desc:    "Article One Description",
			Content: "Article One Content, Article One Content",
		},
		{
			Id:      2,
			Title:   "Article Two",
			Desc:    "Article Two Description",
			Content: "Article Two Content, Article Two Content",
		},
		{
			Id:      3,
			Title:   "Article Three",
			Desc:    "Article Three Description",
			Content: "Article Three Content, Article Three Content",
		},
		{
			Id:      4,
			Title:   "Article Four",
			Desc:    "Article Four Description",
			Content: "Article Four Content, Article Four Content",
		},
		{
			Id:      5,
			Title:   "Article Five",
			Desc:    "Article Five Description",
			Content: "Article Five Content, Article Five Content",
		},
	}
}

func main() {
	fmt.Println("Rest API v1.0")
	parseData()
	handleRequests()
}
