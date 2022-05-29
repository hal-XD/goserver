package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start")
	httpHandler()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func httpHandler() {
	http.HandleFunc("/", test)
	http.HandleFunc("/api/v1/hoge", hoge)
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello!!!! I am go\n")
	fmt.Println("Endpoint hit:test")
}

func hoge(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is api hoge\n")
}
