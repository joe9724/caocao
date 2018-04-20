package main

import (
	"io"
	"net/http"
	"log"
)

func Callback(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello caocao!")
}

func main() {
	http.HandleFunc("/callback", Callback)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}