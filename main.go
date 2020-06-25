package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "up and running...")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/add", addPost).Methods("Post")
	log.Println("Server running on port 8000")
	log.Fatalln(http.ListenAndServe(port, router))
}