package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"` // String => JSON Transfer
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

/*
Not Transfer
{"FirstName":"","LastName":"","Email":"hyun.sang@parkhyunsang.com","CreatedAt":"2021-03-13T16:23:07.319954+09:00"}
String => JSON Transfer
{"first_name":"HyunSang","last_name":"Park","email":"hyun.sang@parkhyunsang.com","created_at":"2021-03-13T16:30:24.447374+09:00"}
*/

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	// 127.0.0.1/bar?name=Hello
	// => Hello Hello
	fmt.Fprint(w, "Hello ", name, "!")
}

func NewHttpHandler() http.Handle {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})
	return mux
}
