package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":3100", router))
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}