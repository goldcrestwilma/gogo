package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	http.HandleFunc("/helloWorld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type helloWorldResponse struct {
	Message string `json:"message"`
}

func helloWorldHandler(rw http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "HelloWorld"}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	callback := r.URL.Query().Get("callback")
	if callback != "" {
		r.Header.Set("Content-Type", "application/javascript")
		fmt.Fprintf(rw, "%s(%s)", callback, string(data))
	} else {
		fmt.Fprint(rw, "Hello World\n")
	}
}

/*
http://localhost:8080/helloWorld?callback=back

back({
	"message": "HelloWorld"
  })
*/
