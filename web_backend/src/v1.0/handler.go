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

// GetUser get user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	model := new(User)
	user := model.Get(1)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&user); err != nil {
		panic(err)
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	user.Username = "hongker"
	user.Email = "test@qq.com"
	user.Password = "123456"
	user.Type = 3
	user.CreatedAt = 1501234567

	result := make(map[string]interface{})
	id := user.Add()

	result["errorNo"] = 0
	result["errorMsg"] = "没有错误"
	result["id"] = id

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}
