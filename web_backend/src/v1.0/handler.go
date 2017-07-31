package main

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

// AddUser 注册
func AddUser(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	r.ParseForm()
	postData, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	result := make(map[string]interface{})

	dat := make(map[string]string)
	if err := json.Unmarshal([]byte(postData), &dat); err != nil {
		panic(err)
	}
	if len(dat["username"]) == 0 || len(dat["email"]) == 0 || len(dat["password"]) == 0 {
		result["errorNo"] = 103
		result["errorMsg"] = "参数不能为空"
		JsonReturn(w, result)
		return
	}

	if user.HasName(dat["username"]) {
		result["errorNo"] = 104
		result["errorMsg"] = "用户名已存在"
		JsonReturn(w, result)
		return
	}

	if user.HasEmail(dat["email"]) {
		result["errorNo"] = 105
		result["errorMsg"] = "邮箱已被使用"
		JsonReturn(w, result)
		return
	}

	user.Username = dat["username"]
	user.Email = dat["email"]
	user.Password = dat["password"]
	user.Type = 3
	user.CreatedAt = time.Now().Unix()

	user.Add()

	result["errorNo"] = 0
	result["errorMsg"] = "没有错误"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

// Login 登录
func Login(w http.ResponseWriter, r *http.Request) {
	model := new(User)
	r.ParseForm()
	postData, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	dat := make(map[string]string)
	if err := json.Unmarshal([]byte(postData), &dat); err != nil {
		panic(err)
	}

	username := dat["username"]
	password := dat["password"]

	result := make(map[string]interface{})

	if len(username) == 0 || len(password) == 0 {
		result["errorNo"] = 103
		result["errorMsg"] = "参数不能为空"
		JsonReturn(w, result)
		return
	}
	user, err := model.Auth(username, password)

	if err != nil {
		result["errorNo"] = err.No
		result["errorMsg"] = err.Msg
	}

	result["data"] = &user

	JsonReturn(w, result)
	return
}

// JsonReturn 输出json
func JsonReturn(w http.ResponseWriter, result map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}
