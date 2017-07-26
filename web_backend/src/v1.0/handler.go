package main

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
)

// Index function
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// Add test post
func Add(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Post can see!")
}

// ArticleIndex handle article get
func ArticleIndex(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			1,
			"this is article title",
			"this is article content",
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		panic(err)
	}
}
