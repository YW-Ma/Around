package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("started-service")
	http.HandleFunc("/upload", uploadHandler) // handler的是/upload的请求(各METHOD靠HTTP Router来分派handler)
	http.HandleFunc("/search", searchHandler)
	log.Fatal(http.ListenAndServe(":8080", nil)) // 默认的这个router，没有分派的功能，所有的METHOD都调用同一个handler。
}
