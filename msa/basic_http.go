package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"
)

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	// 출력 필드를 "message"로 변경
	Message string `json:"message"`

	// 이 필드는 출력하지 않음
	Author string `json:"-"`

	// 값이 비어 있으면 필드를 출력하지 않음
	Date string `json:", omitempty"`

	// 출력을 문자열로 변환하고 이름을 "id"로 바꾼다.
	Id int `json:"id, string"`
}

func main() {
	port := 8080

	//http.HandleFunc("/helloWorld", helloWorldHandler)
	handler := newValidationHandler(newHelloWorldHandler())

	catHandler := http.FileServer(http.Dir("./images"))
	http.Handle("/cat/", http.StripPrefix("/cat/", catHandler))

	http.Handle("/helloWorld", handler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type validationHandler struct {
	next http.Handler
}

type validationContextKey string

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return
	}

	c := context.WithValue(r.Context(), validationContextKey("name"), request.Name)
	r = r.WithContext(c)

	h.next.ServeHTTP(rw, r)
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

type helloWorldHandler struct{}

func newHelloWorldHandler() http.Handler {
	return helloWorldHandler{}
}

func (h helloWorldHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	name := r.Context().Value(validationContextKey("name")).(string)
	response := helloWorldResponse{Message: "Hello " + name}

	encdoer := json.NewEncoder(rw)
	encdoer.Encode(response)
}

func fetchGoogle(t *testing.T) {
	r, _ := http.NewRequest("GET", "https://google.com", nil)

	timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1*time.Millisecond)
	defer cancelFunc()

	r = r.WithContext(timeoutRequest)

	_, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
