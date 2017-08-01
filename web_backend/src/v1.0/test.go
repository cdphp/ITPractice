package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"v1.0/vendor"
)

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method: ", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("src/html/login.gtpl")
		if err != nil {
			fmt.Println(err)
			return
		}
		t.Execute(w, nil)

	} else {
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}
func login2(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	r.ParseForm()
	if r.Method == "GET" {
		t, _ := template.ParseFiles("src/html/login.gtpl")
		//w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		sess.Set("username", r.Form["username"])
		http.Redirect(w, r, "/", 302)
	}
}

func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)
	createtime := sess.Get("createtime")
	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, r)
		sess = globalSessions.SessionStart(w, r)
	}
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}
func sayHelloName(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)

	sess.Set("createtime", time.Now().Unix())
	fmt.Fprintf(w, "set session")

}
func getSession(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)

	sess.Set("createtime", time.Now().Unix())
	fmt.Println(sess.Get("createtime"))
}
func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/get", getSession)
	http.HandleFunc("/login", login)
	http.HandleFunc("/login2", login2)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalf("Listen and server", err)
	}
}

var globalSessions *vendor.Manager

func init() {
	globalSessions, _ = vendor.NewSessionManager("memory", "goSessionid", 3600)
	go globalSessions.GC()
	fmt.Println("fd")
}
