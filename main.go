package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("started-service")
	// http.HandleFunc("/upload", uploadHandler) // handler的是/upload的请求(各METHOD靠HTTP Router来分派handler)
	// http.HandleFunc("/search", searchHandler)
	r := mux.NewRouter()
	r.Handle("/upload", http.HandlerFunc(uploadHandler)).Methods("POST", "OPTIONS") 
	r.Handle("/search", http.HandlerFunc(searchHandler)).Methods("GET", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8080", r)) // 如果不用r，用nil则用默认的router，没有分派的功能，所有的METHOD都调用同一个handler。
}
