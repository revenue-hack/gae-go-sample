package main

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/hoge", handle2)
	appengine.Main()
}

func handle(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Response{Status: "ok", Message: "Hello world."})
}

func handle2(w http.ResponseWriter, r *http.Request) {
	type userS struct {
		Id int
		Name string
	}
	user := &userS{Id: 1,Name: "hogehgoehgoheoghe"}
	j, _ := json.Marshal(user)
	json.NewEncoder(w).Encode(Response{Status: "OKOKOKO", Message: string(j)})
}
