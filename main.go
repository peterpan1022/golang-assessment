package main

import (
	"io"
	"log"
	"net/http"
)

func FindHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"foo": "bar"}`)
}

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	io.WriteString(w, `{"bar": "foo"}`)
}

func main() {
	http.HandleFunc("/find", FindHandler)
	http.HandleFunc("/compare", CompareHandler)
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
